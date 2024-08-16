package Anubis

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	rr, _ := Run("hello.py")
	out, _ := io.ReadAll(rr.StdOut)
	assert.Equal(t, string(out), "Hello World\n", "Expcted: %d, Got: %d", "Hello World", string(out))
}
