package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

// -- Helper Functions --

// Helper function for testing renderPage.
func DoRenderTest(input, expected string, t *testing.T) {

	inputReadCloser := ioutil.NopCloser(strings.NewReader(input))

	output := renderPage(inputReadCloser)

	if output != expected {
		t.Error("Expected ", expected, ", got ", output)
	}
}

// -- Tests for renderPage --

func TestRenderPageTitle(t *testing.T) {
	input := "# Title"
	expected := "\033[1mTitle\033[0m\n"

	DoRenderTest(input, expected, t)
}

func TestRenderPageDescription(t *testing.T) {
	input := "> Description"
	expected := "Description\n"

	DoRenderTest(input, expected, t)
}

func TestRenderPageExampleHeader(t *testing.T) {
	input := "- Example Header"
	expected := "-\033[32;1m Example Header\033[0m\n"

	DoRenderTest(input, expected, t)
}

func TestRenderPageExampleCode(t *testing.T) {
	input := "` Example Code `"
	expected := "  \033[31;1m Example Code \033[0m\n"

	DoRenderTest(input, expected, t)
}

func TestRenderPageExampleCodePlaceHolder(t *testing.T) {
	input := "` Example {{ placeholder }} Code `"
	expected := "  \033[31;1m Example \033[0m placeholder \033[31;1m Code \033[0m\n"

	DoRenderTest(input, expected, t)
}

// -- Tests for queryGithub --
//func TestQueryGithub(t *testing.T) {
//}

// -- Tests for getPage --
func TestGetPageGood(t *testing.T) {
	cmd := "grep"

	page, err := getPage(cmd)

	if err != nil {
		t.Error("Got Error for know good page")
	}

	page.Close()
}

func TestGetPageBad(t *testing.T) {
	cmd := "not_a_real_command"

	_, err := getPage(cmd)

	if err == nil {
		t.Error("Didn't get an Error for a bad page")
	}
}
