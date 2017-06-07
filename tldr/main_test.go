package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

// Tests for getPlatform
func TestGetPaltform(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic",
			input:    "linux",
			expected: "linux",
		},
		{
			name:     "Basic",
			input:    "darwin",
			expected: "osx",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			osPlatform = tc.input

			output := getPlatform()
			if output != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, output)
			}
		})
	}
}

// Tests for renderPage --
func TestRenderPage(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Title",
			input:    "# Title",
			expected: "\033[1mTitle\033[0m\n",
		},
		{
			name:     "Description",
			input:    "> Description",
			expected: "Description\n",
		},
		{
			name:     "Header",
			input:    "- Example Header",
			expected: "-\033[32;1m Example Header\033[0m\n",
		},
		{
			name:     "Code",
			input:    "` Example Code `",
			expected: "  \033[31;1m Example Code \033[0m\n",
		},
		{
			name:     "Placeholder",
			input:    "` Example {{ placeholder }} Code `",
			expected: "  \033[31;1m Example \033[0m placeholder \033[31;1m Code \033[0m\n",
		},
		{
			name:     "Empty Line",
			input:    "\n",
			expected: "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			p := Page{"dummy", ioutil.NopCloser(strings.NewReader(tc.input))}

			output := p.Render()

			if output != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, output)
			}
		})
	}
}

// -- Tests for queryGithub --
//func TestQueryGithub(t *testing.T) {
//}

// -- Tests for getPage --
func TestGetPageGood(t *testing.T) {
	page := Page{cmd: "grep"}

	err := page.Fetch()

	if err != nil {
		t.Error("Got Error for know good page", err)
	}

	page.Close()
}

func TestGetPageBad(t *testing.T) {
	page := Page{cmd: "not_a_real_command"}

	err := page.Fetch()

	if err == nil {
		t.Error("Didn't get an Error for a bad page")
	}
}
