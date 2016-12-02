package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// -- Constants --

const (
	// Template URL for TLDR pages
	urlBase = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages/%s/%s.md"

	// ASCII codes for colourful output
	escapePrefix = "\033["
	escapeSuffix = "m"
	bright       = ";1"

	reset = escapePrefix + "0" + escapeSuffix
	bold  = escapePrefix + "1" + escapeSuffix

	red   = escapePrefix + "31" + bright + escapeSuffix
	green = escapePrefix + "32" + bright + escapeSuffix
)

// -- Functions --

func getPlatform() string {
	platform := runtime.GOOS

	if platform == "darwin" {
		platform = "osx"
	}

	return platform
}

func queryGithub(cmd, platform string) io.ReadCloser {

	url := fmt.Sprintf(urlBase, platform, cmd)

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode == 200 {
		return resp.Body
	}

	return nil
}

// -- Page --

type Page struct {
	cmd     string
	content io.ReadCloser
}

func (page *Page) Close() {
	page.content.Close()
}

func (page *Page) Fetch() error {

	// First try to get the platform page
	body := queryGithub(page.cmd, getPlatform())
	if body != nil {
		page.content = body
		return nil
	}

	// If that fails try to get the common page
	body = queryGithub(page.cmd, "common")
	if body != nil {
		page.content = body
		return nil
	}

	// If we got something return it, if not bail
	return errors.New("Unable to find page for " + page.cmd)
}

func (page *Page) Render() string {
	scanner := bufio.NewScanner(page.content)

	rendered := ""

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "#"):
			// Heading
			rendered += bold + strings.TrimLeft(line, "# ") + reset + "\n"

		case strings.HasPrefix(line, ">"):
			// Description
			rendered += strings.TrimLeft(line, "> ") + "\n"

		case strings.HasPrefix(line, "-"):
			// Example header
			rendered += "-" + green + strings.TrimLeft(line, "-") + reset + "\n"

		case strings.HasPrefix(line, "`"):
			// Example code
			line = strings.Replace(line, "{{", reset, -1)
			line = strings.Replace(line, "}}", red, -1)
			rendered += "  " + red + strings.Trim(line, "`") + reset + "\n"

		default:
			// Probably a blank line
			rendered += line + "\n"
		}
	}

	return rendered
}

// -- Main --

func main() {
	// TODO - Document this and everythng else
	// TODO - Add in a verbose flag for more verbose output (with error handling)
	// TODO - Better arg parsing
	if len(os.Args) < 2 {
		fmt.Println("Need a page to look up!")
		os.Exit(1)
	}
	cmd := os.Args[1]

	// TODO - Cache results (and add flag to clear/update cache)
	page := Page{cmd: cmd}
	err := page.Fetch()

	// TODO - Better error handling (everywhere not just here)
	if err != nil {
		fmt.Println("Oh shit something bad happened!")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(page.Render())
	page.Close()
}
