package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

// -- Helper Functions --

// Helper function for testing renderPage.
func DoRenderTest(t *testing.T, input, expected string) {

	p := Page{"dummy", ioutil.NopCloser(strings.NewReader(input))}

	output := p.Render()

	if output != expected {
		t.Error("Expected ", expected, ", got ", output)
	}
}

// -- Tests for renderPage --

func TestRenderPageTitle(t *testing.T) {
	input := "# Title"
	expected := "\033[1mTitle\033[0m\n"

	DoRenderTest(t, input, expected)
}

func TestRenderPageDescription(t *testing.T) {
	input := "> Description"
	expected := "Description\n"

	DoRenderTest(t, input, expected)
}

func TestRenderPageExampleHeader(t *testing.T) {
	input := "- Example Header"
	expected := "-\033[32;1m Example Header\033[0m\n"

	DoRenderTest(t, input, expected)
}

func TestRenderPageExampleCode(t *testing.T) {
	input := "` Example Code `"
	expected := "  \033[31;1m Example Code \033[0m\n"

	DoRenderTest(t, input, expected)
}

func TestRenderPageExampleCodePlaceHolder(t *testing.T) {
	input := "` Example {{ placeholder }} Code `"
	expected := "  \033[31;1m Example \033[0m placeholder \033[31;1m Code \033[0m\n"

	DoRenderTest(t, input, expected)
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
