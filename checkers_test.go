package Anubis

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckCasePass(t *testing.T) {
	expected := bytes.NewBuffer([]byte("1\n2\n3\n4\n5"))
	actual := bytes.NewBuffer([]byte("1\n2\n3\n4\n5"))

	pass := checkCase(actual, expected)
	require.True(t, pass)

	expected = bytes.NewBuffer([]byte(
		`
		1 2
		3 5  
		5 0

		4 8`))
	actual = bytes.NewBuffer([]byte(
		`
		1 2
		3 5  
			5 0

		4 8`))

	pass = checkCase(actual, expected)
	require.True(t, pass)
}
