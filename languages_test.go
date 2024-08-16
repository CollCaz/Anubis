package Anubis

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProgLang(t *testing.T) {
	t.Parallel()
	for ext, actualLang := range supportedFileTypes {
		file := "main." + ext
		lang, err := GetProgLang(file)
		require.Nil(t, err)
		require.Equal(t, actualLang.Name, lang.Name, "Expected: %d, Got: %d", actualLang.Name, lang.Name)
	}
}

func TestGetProgLangErr(t *testing.T) {
	t.Parallel()
	lang, err := GetProgLang("main.Fake")
	require.Equal(t, lang, ProgLang{}, "Expected: %d, Got: %d", 0, lang)
	require.NotNil(t, err)
}

func TestAddProgLang(t *testing.T) {
	t.Parallel()
	javaRunner := func(codeFile string, commandRunner CommandRunner) (RunOutput, error) {
		rr := RunOutput{
			ExitStatus: 0,
			StdOut:     bytes.NewBuffer([]byte(codeFile)),
		}
		return rr, nil
	}
	AddProgLang("Java", "java", javaRunner)

	fileName := "HelloWorld.java"
	rr, err := Run(fileName, &LocalCmdRunner{})
	require.Nil(t, err)
	stdout, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err)
	require.Equal(t, string(stdout), fileName)
}
