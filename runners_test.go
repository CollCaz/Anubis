package Anubis

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalCommandRunner(t *testing.T) {
	t.Parallel()
	fileName := fmt.Sprintf("%s/test.txt", t.TempDir())
	file, err := os.Create(fileName)
	defer os.Remove(fileName)
	require.Nil(t, err)
	file.WriteString("hello world")
	file.Close()

	f, err := os.Open(fileName)
	require.Nil(t, err)

	lcr := LocalCmdRunner{Input: f}
	cmd := exec.Command("cat")
	rr, err := lcr.RunCommand(cmd)
	require.Nil(t, err)
	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err)
	require.Equal(t, "hello world", string(out))
}

func TestLocalNsjailRunner(t *testing.T) {
	t.Parallel()
	dir := os.TempDir()
	fileName := "test.txt"

	file, err := os.CreateTemp(dir, fileName)
	require.Nil(t, err, err)
	file.WriteString("hello world")
	file.Close()

	lnr := LocalNsjailRunner{Input: file}
	cmd := exec.Command("/home/coll/a.out")
	rr, err := lnr.RunCommand(cmd)
	require.Nil(t, err, err)
	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err, err)
	require.Equal(t, "hello world", string(out), cmd.String())
}

func TestLocalRunner(t *testing.T) {
	t.Parallel()
	filename := fmt.Sprintf("%s/AnubisRunnerTestHello.py", t.TempDir())
	pyFile, err := os.Create(filename)
	require.Nil(t, err)
	_, err = pyFile.WriteString(`print(int(input())+20)`)
	pyFile.Close()
	defer os.Remove(filename)
	require.Nil(t, err)

	inputFileName := fmt.Sprintf("%s/AnubisRunnerTestHelloInput.txt", t.TempDir())
	inputFile, err := os.Create(inputFileName)
	require.Nil(t, err)
	_, err = inputFile.WriteString("400")
	inputFile.Close()
	defer os.Remove(inputFileName)
	require.Nil(t, err)

	in, err := os.Open(inputFileName)
	rr, err := Run(pyFile, &LocalCmdRunner{Input: in}, slog.New(&noopLogHandler{}))
	require.Nil(t, err)

	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err)

	require.Equal(t, "420\n", string(out), "Expcted:\n%s,\nGot:\n%s", "420", string(out))
}

func TestPythonRunner(t *testing.T) {
	t.Parallel()
	codeFile, _ := os.CreateTemp(os.TempDir(), "code.py")
	require.NotNil(t, codeFile)
	_, _ = codeFile.WriteString("print('hello world')")
	lcr := &LocalCmdRunner{}
	require.NotNil(t, lcr)
	rr, err := PythonRunner(codeFile, lcr)
	require.Nil(t, err)
	bytesFail, _ := io.ReadAll(rr.StdErr)
	require.Empty(t, bytesFail)
	bytes, _ := io.ReadAll(rr.StdOut)
	require.Equal(t, "hello world\n", string(bytes))
}
