package Anubis

import (
	"bytes"
	"io"
	"os/exec"
)

type RunResult struct {
	ExitStatus int
	StdOut     io.Reader
	StdErr     io.Reader
}

type CodeRunner func(codeFile string) RunResult

func Run(codeFile string) (RunResult, error) {
	progLang, err := GetProgLang(codeFile)
	if err != nil {
		return RunResult{}, err
	}
	return progLang.Runner(codeFile), nil
}

func PythonRunner(codeFile string) RunResult {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	app := "python3"
	command := exec.Command(app, codeFile)
	command.Stdout = &stdout
	command.Stderr = &stderr
	rr := RunResult{}
	_ = command.Run()
	rr.StdOut = &stdout
	rr.StdErr = &stderr
	return rr
}
