package wordlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEFF_Long(t *testing.T) {
	assert.NotEmpty(t, EFF_Long())
}
