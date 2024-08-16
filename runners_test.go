package Anubis

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	filename := "AnubisRunnerTestHello.py"
	pyFile, err := os.Create(filename)
	require.Nil(t, err)
	_, err = pyFile.WriteString(`print(int(input())+20)`)
	defer pyFile.Close()
	defer os.Remove(filename)
	require.Nil(t, err)

	inputFileName := "AnubisRunnerTestHelloInput.txt"
	inputFile, err := os.Create(inputFileName)
	require.Nil(t, err)
	_, err = inputFile.WriteString("400")
	defer inputFile.Close()
	defer os.Remove(inputFileName)
	require.Nil(t, err)

	in, err := os.Open(inputFileName)
	rr, err := Run(filename, &LocalCmdRunner{Input: in})
	fmt.Println(rr.StdErr)
	require.Nil(t, err)

	out, err := io.ReadAll(rr.StdOut)
	require.Nil(t, err)

	require.Equal(t, "420\n", string(out), "Expcted:\n%s,\nGot:\n%s", "420", string(out))
}
