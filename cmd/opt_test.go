package cmd

import (
	"testing"

	"github.com/muesli/termenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithMarkdown(t *testing.T) {
	root := New(WithMarkdown())
	assert.NotContains(t, root.Long, termenv.OSC)

	prof, _, err := root.Find([]string{"profiles"})
	require.NoError(t, err)
	assert.Equal(t, "markdown", prof.Annotations["format"])
}

func TestWithRaw(t *testing.T) {
	assert.NotContains(t, New(WithMarkdown()).Long, termenv.OSC)
}

func TestWithVersion(t *testing.T) {
	cmd := New(WithVersion("1.0.0"))
	assert.NotEmpty(t, cmd.Version)
}
