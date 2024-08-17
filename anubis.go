package Anubis

import (
	"context"
	"log/slog"
)

type Anubis struct {
	Logger        *slog.Logger
	CommandRunner CommandRunner
}

type AnubisConfig struct {
	Logger        *slog.Logger
	CommandRunner CommandRunner
}

func NewAnubis(ac AnubisConfig) Anubis {
	a := Anubis{}
	if ac.Logger == nil {
		ac.Logger = &slog.Logger{}
	}
	if ac.CommandRunner == nil {
		ac.CommandRunner = &LocalCmdRunner{}
	}

	a.Logger = ac.Logger
	a.CommandRunner = ac.CommandRunner

	return a
}

func (a *Anubis) NewSubmission(codeFilePath string, testCases TestCases) Submission {
	if a.Logger == nil {
		a.Logger = slog.New(&noopLogHandler{})
	}
	so := Submission{
		CodeFile:      codeFilePath,
		TestCases:     testCases,
		CommandRunner: a.CommandRunner,
		Logger:        a.Logger,
	}

	return so
}

type noopLogHandler struct{}

func (n *noopLogHandler) Enabled(_ context.Context, l slog.Level) bool {
	return false
}

func (n *noopLogHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (n *noopLogHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return &noopLogHandler{}
}

func (n *noopLogHandler) WithGroup(_ string) slog.Handler {
	return &noopLogHandler{}
}
