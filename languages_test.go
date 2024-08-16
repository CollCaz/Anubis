package Anubis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProgLang(t *testing.T) {
	t.Parallel()
	for ext, actualLang := range supportedFileTypes {
		file := "main." + ext
		lang, err := GetProgLang(file)
		assert.Nil(t, err)
		assert.Equal(t, actualLang, lang, "Expected: %d, Got: %d", actualLang, lang)
	}
}

func TestGetProgLangErr(t *testing.T) {
	t.Parallel()
	lang, err := GetProgLang("main.Fake")
	assert.Equal(t, lang, ProgLang{}, "Expected: %d, Got: %d", 0, lang)
	assert.NotNil(t, err)
}
