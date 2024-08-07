package rand

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASCII(t *testing.T) {
	tests := []struct {
		n int
	}{
		{0},
		{1},
		{10},
		{100},
		{1000},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Len(t, ASCII(tt.n), tt.n)
		})
	}
}

func TestAlpha(t *testing.T) {
	tests := []struct {
		n int
	}{
		{0},
		{1},
		{10},
		{100},
		{1000},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Len(t, Alpha(tt.n), tt.n)
		})
	}
}

func TestAlphaNum(t *testing.T) {
	tests := []struct {
		n int
	}{
		{0}, {1}, {10}, {100}, {1000},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Len(t, AlphaNum(tt.n), tt.n)
		})
	}
}

func TestNumeric(t *testing.T) {
	tests := []struct {
		n int
	}{
		{0},
		{1},
		{10},
		{100},
		{1000},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			assert.Len(t, Numeric(tt.n), tt.n)
		})
	}
}

func TestShuffleSlice(t *testing.T) {
	got := make([]int, 100)
	for i := range got {
		got[i] = i
	}
	expected := slices.Clone(got)
	ShuffleSlice(got)
	assert.NotEqual(t, expected, got)
}
