package Anubis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
)

type CommandRunner interface {
	RunCommand(command *exec.Cmd) (RunOutput, error)
	SetInput(file *os.File)
}

type LocalCmdRunner struct {
	Input *os.File
}

func (lcr *LocalCmdRunner) SetInput(file *os.File) {
	lcr.Input = file
}

func (lcr *LocalCmdRunner) RunCommand(command *exec.Cmd) (RunOutput, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	rr := RunOutput{
		ExitStatus: 0,
		StdOut:     &stdout,
		StdErr:     &stderr,
	}

	command.Stdin = lcr.Input
	defer lcr.Input.Close()
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			rr.ExitStatus = command.ProcessState.ExitCode()
			return rr, err
		}
		return rr, err
	}
	command.Wait()

	return rr, nil
}

type RunOutput struct {
	ExitStatus int
	StdOut     io.Reader
	StdErr     io.Reader
}

func (rr *RunOutput) String() string {
	outString := "Error Reading StdOut"
	outBytes, err := io.ReadAll(rr.StdOut)
	if err == nil {
		outString = string(outBytes)
	}
	errString := "Error Reading StdErr"
	errBytes, err := io.ReadAll(rr.StdErr)
	if err == nil {
		errString = string(errBytes)
	}

	s := fmt.Sprintf("Exit Status: %d\nStdOut: %s\nStdErr:= %s", rr.ExitStatus, outString, errString)

	return s
}

type CodeRunner func(codeFile string, commandRunner CommandRunner) (RunOutput, error)

func Run(codeFile string, commandRunner CommandRunner, logger *slog.Logger) (RunOutput, error) {
	progLang, err := GetProgLang(codeFile)
	if err != nil {
		logger.Error(err.Error())
		return RunOutput{}, err
	}
	return progLang.Runner(codeFile, commandRunner)
}

func PythonRunner(codeFile string, commandRunner CommandRunner) (RunOutput, error) {
	app := "python3"
	command := exec.Command(app, codeFile)
	rr, err := commandRunner.RunCommand(command)
	return rr, err
}
