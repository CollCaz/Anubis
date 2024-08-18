package Anubis

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// creates a submission struct for the problem of
// multiplying the inputs by 2
func createSubmission(t *testing.T) (Submission, func()) {
	testCases := make(TestCases)
	for i := 0; i < 10; i++ {
		fileIn := fmt.Sprintf("%sin%d", t.TempDir(), i)
		inFile, err := os.Create(fileIn)
		require.Nil(t, err)
		fileOut := fmt.Sprintf("%sout%d", t.TempDir(), i)
		outFile, err := os.Create(fileOut)
		require.Nil(t, err)
		os.Create(fileOut)
		testCases[fileIn] = fileOut

		inFile.WriteString(fmt.Sprintf("%d %d %d %d", i, i+1, i+2, i+3))
		outFile.WriteString(fmt.Sprintf("%d %d %d %d", i*2, (i+1)*2, (i+2)*2, (i+3)*2))
	}

	submission := Submission{
		TestCases:     testCases,
		CommandRunner: &LocalCmdRunner{},
		Logger:        slog.New(&noopLogHandler{}),
	}

	clean := func() {
		for in, out := range testCases {
			os.Remove(in)
			os.Remove(out)
		}
	}

	return submission, clean
}

func TestCheckAllPass(t *testing.T) {
	t.Parallel()
	codeName := fmt.Sprintf("%s/TestCheckAllPass.py", t.TempDir())
	code, err := os.Create(codeName)
	defer os.Remove(codeName)
	code.WriteString(fmt.Sprintf(
		"inp=[int(x) * 2 for x in input().split()]\n%s",
		`for x in inp:print(x, end=" ")`))
	code.Close()
	require.Nil(t, err)
	sub, clean := createSubmission(t)
	sub.CodeFile = codeName
	subOut, err := sub.CheckAll()
	require.Nil(t, err)
	require.Equal(t, AC, subOut.Status)

	t.Cleanup(clean)
}

func TestCheckAllFail(t *testing.T) {
	t.Parallel()
	codeName := fmt.Sprintf("%s/TestCheckAllFail.py", t.TempDir())
	code, err := os.Create(codeName)
	defer os.Remove(codeName)
	code.WriteString(fmt.Sprintf(
		"inp=[int(x) * 100 for x in input().split()]\n%s",
		`for x in inp:print(x, end=" ")`))
	code.Close()
	require.Nil(t, err, err)
	sub, clean := createSubmission(t)
	sub.CodeFile = codeName
	subOut, err := sub.CheckAll()
	require.Nil(t, err, err)
	require.Equal(t, Failed, subOut.Status)

	t.Cleanup(clean)
}

func TestCheckCasePass(t *testing.T) {
	t.Parallel()
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

func TestCheckCaseFail(t *testing.T) {
	t.Parallel()
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
	require.False(t, pass)
}
