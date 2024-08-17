package Anubis

import "log"

type Anubis struct {
	logger        *log.Logger
	CommandRunner CommandRunner
}

type AnubisConfig struct {
	logger        *log.Logger
	CommandRunner CommandRunner
}

func NewAnubis(ac AnubisConfig) Anubis {
	a := Anubis{}
	if ac.logger == nil {
		ac.logger = &log.Logger{}
	}
	if ac.CommandRunner == nil {
		ac.CommandRunner = &LocalCmdRunner{}
	}

	a.logger = ac.logger
	a.CommandRunner = ac.CommandRunner

	return a
}

func (a Anubis) NewSubmission(codeFilePath string, testCases TestCases) Submission {
	so := Submission{
		CodeFile:      codeFilePath,
		TestCases:     testCases,
		CommandRunner: a.CommandRunner,
	}

	return so
}
