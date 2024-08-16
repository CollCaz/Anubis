package Anubis

import (
	"errors"
	"strings"
)

// Programming Languages
type ProgLang struct {
	Name   string
	Runner CodeRunner
}

// file extention -> ProgLang
var supportedFileTypes = map[string]ProgLang{
	"c":   {Name: "C"},
	"cpp": {Name: "C++"},
	"py":  {Name: "Python"},
}

func GetProgLang(filePath string) (ProgLang, error) {
	filePath = strings.ToLower(filePath)
	tmp := strings.Split(filePath, ".")
	ext := tmp[len(tmp)-1]

	prog, ok := supportedFileTypes[ext]
	if !ok {
		return ProgLang{}, errors.New("Language not supported")
	}

	return prog, nil
}
