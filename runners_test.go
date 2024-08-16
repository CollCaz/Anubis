package Anubis

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	filename := "AnubisRunnerTestHello.py"
	pyFile, err := os.Create(filename)
	assert.Nil(t, err)
	defer os.Remove(pyFile.Name())
	_, err = pyFile.WriteString(`print("Hello World")`)
	assert.Nil(t, err)

	rr, _ := Run(filename)
	out, _ := io.ReadAll(rr.StdOut)
	assert.Equal(t, string(out), "Hello World\n", "Expcted: %d, Got: %d", "Hello World", string(out))
}
