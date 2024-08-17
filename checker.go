package Anubis

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// maps test input files with their expected output files
type TestCases map[string]string

type Submission struct {
	CodeFile      string
	TestCases     TestCases
	CommandRunner CommandRunner
	Logger        *slog.Logger
}

type SubmissionStatus int

const (
	AC SubmissionStatus = iota
	CompileError
	RunTimeError
	Failed
)

type SubmissionOut struct {
	Status string
	// 0 if no failure, otherwise the first failed test case
	FailedOn int
	StdOut   io.Reader
	StdErr   io.Reader
}

func (s *Submission) CheckAll() (SubmissionOut, error) {
	currentTest := 0
	for in, out := range s.TestCases {
		currentTest++
		s.Logger.Info(fmt.Sprintf("Staring test %d", currentTest))
		inFile, err := os.Open(in)
		defer inFile.Close()

		if err != nil {
			s.Logger.Error(fmt.Sprintf("Error opening the input file for test %d: %s", currentTest, err.Error()))
			return SubmissionOut{Status: "FailedOpeningInput", FailedOn: currentTest}, err
		}

		outFile, err := os.Open(out)
		defer outFile.Close()

		if err != nil {
			return SubmissionOut{Status: "FailedOpeningOutput", FailedOn: currentTest}, err
		}

		s.CommandRunner.SetInput(inFile)
		rr, err := Run(s.CodeFile, s.CommandRunner)
		if err != nil {
			return SubmissionOut{Status: "FailedRunningCode", FailedOn: currentTest, StdOut: rr.StdOut, StdErr: rr.StdErr}, err
		}
		actualOutput := rr.StdOut
		if !checkCase(outFile, actualOutput) {
			return SubmissionOut{Status: "Failed", FailedOn: currentTest, StdOut: rr.StdOut, StdErr: rr.StdErr}, err
		}
	}
	so := SubmissionOut{Status: "AC"}
	return so, nil
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
