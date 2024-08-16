package Anubis

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProgLang(t *testing.T) {
	t.Parallel()
	for ext, actualLang := range supportedFileTypes {
		file := "main." + ext
		lang, err := GetProgLang(file)
		assert.Nil(t, err)
		assert.Equal(t, actualLang.Name, lang.Name, "Expected: %d, Got: %d", actualLang.Name, lang.Name)
	}
}

func TestGetProgLangErr(t *testing.T) {
	t.Parallel()
	lang, err := GetProgLang("main.Fake")
	assert.Equal(t, lang, ProgLang{}, "Expected: %d, Got: %d", 0, lang)
	assert.NotNil(t, err)
}

func TestAddProgLang(t *testing.T) {
	t.Parallel()
	javaRunner := func(codeFile string) (RunResult, error) {
		rr := RunResult{
			ExitStatus: 0,
			StdOut:     bytes.NewBuffer([]byte(codeFile)),
		}
		return rr, nil
	}
	AddProgLang("Java", "java", javaRunner)

	fileName := "HelloWorld.java"
	rr, err := Run(fileName)
	assert.Nil(t, err)
	stdout, err := io.ReadAll(rr.StdOut)
	assert.Nil(t, err)
	assert.Equal(t, string(stdout), fileName)
}
