package profiles

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helpFunc(t *testing.T) {
	cmd := New()

	var raw strings.Builder
	cmd.SetOut(&raw)
	helpFunc(cmd, nil)
	assert.NotZero(t, raw.Len())

	var markdown strings.Builder
	cmd.SetOut(&markdown)
	SetMarkdown(cmd, true)
	helpFunc(cmd, nil)
	assert.NotZero(t, markdown.Len())

	assert.NotEqual(t, raw.String(), markdown.String())
}
