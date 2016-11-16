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
	urlBase = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages/"

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

	url := urlBase + "/" + platform + "/" + cmd + ".md"

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	if resp.StatusCode == 200 {
		return resp.Body
	}

	return nil
}

func getPage(cmd string) (io.ReadCloser, error) {

	// First try to get the platform page
	platform := getPlatform()
	page := queryGithub(cmd, platform)
	if page != nil {
		return page, nil
	}

	// If that fails try to get the common page
	page = queryGithub(cmd, "common")
	if page != nil {
		return page, nil
	}

	// If that doesn't work return an error
	return nil, errors.New("Unable to find page for " + cmd)
}

func renderPage(page io.ReadCloser) string {
	scanner := bufio.NewScanner(page)

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
	page, err := getPage(cmd)

	// TODO - Better error handling (everywhere not just here)
	if err != nil {
		fmt.Println("Oh shit something bad happened!")
		//fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer page.Close()

	rendered := renderPage(page)
	fmt.Println(rendered)
}
