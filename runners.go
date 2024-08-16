package Anubis

import (
	"io"
)

type RunResult struct {
	ExitStatus int
	StdOut     io.Reader
	StdErr     io.Reader
}

type CodeRunner func(codeFile string) RunResult

func Run(codeFile string, cr CodeRunner) RunResult {
	return cr(codeFile)
}
