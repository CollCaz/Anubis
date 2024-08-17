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

func TestLocalIsolateRunner(t *testing.T) {
	t.Parallel()
	//fixme: find a better way of giving isolate
	// the full path to files
	dir, err := os.UserHomeDir()
	require.Nil(t, err, err)
	fileName := fmt.Sprintf("%s/test.txt", dir)

	fmt.Println("FIX THIS")
	t.FailNow()

	file, err := os.Create(fileName)
	defer os.Remove(fileName)
	require.Nil(t, err, err)
	file.WriteString("hello world")
	file.Close()

	f, err := os.Open(fileName)
	require.Nil(t, err, err)

	lir := LocalIsolateRunner{Input: f}
	cmd := exec.Command("/bin/cat")
	rr, err := lir.RunCommand(cmd)
	fmt.Println(rr.StdErr)
	require.Nil(t, err, err)
	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err, err)
	fmt.Println(rr.StdErr)
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
	rr, err := Run(filename, &LocalCmdRunner{Input: in}, slog.New(&noopLogHandler{}))
	fmt.Println(rr.StdErr)
	require.Nil(t, err)

	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err)

	require.Equal(t, "420\n", string(out), "Expcted:\n%s,\nGot:\n%s", "420", string(out))
}
