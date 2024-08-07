package wordlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEFF_Short2(t *testing.T) {
	assert.NotEmpty(t, EFF_Short2())
}
