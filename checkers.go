package Anubis

import (
	"bufio"
	"io"
	"strings"
)

type CheckerOutput struct {
	TestsFailed []int
}

type TestCases map[string]string

type Checker struct {
	TestCases TestCases
}

func checkCase(actual, expected io.Reader) bool {
	actualScanner := bufio.NewScanner(actual)
	expectedScanner := bufio.NewScanner(expected)

	for expectedScanner.Scan() {
		expectedLine := strings.TrimSpace(expectedScanner.Text())
		actualScanner.Scan()
		actualLine := strings.TrimSpace(actualScanner.Text())
		if expectedLine != actualLine {
			return false
		}
	}
	if actualScanner.Scan() {
		return false
	}

	return true
}
