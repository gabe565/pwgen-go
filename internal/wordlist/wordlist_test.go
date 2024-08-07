package wordlist

import (
	"strconv"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestWordlist_RandWords(t *testing.T) {
	wl := EFF_Long()

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
			assert.Len(t, wl.RandWords(tt.n), tt.n)
		})
	}
}

func TestWordlist_RandWordsWithNumber(t *testing.T) {
	wl := EFF_Long()

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
			got := wl.RandWordsWithNumber(tt.n)
			assert.Len(t, got, tt.n)
			var numCount int
			for _, word := range got {
				for _, r := range word {
					if unicode.IsDigit(r) {
						numCount++
					}
				}
			}
			assert.Equal(t, min(1, tt.n), numCount)
		})
	}
}
