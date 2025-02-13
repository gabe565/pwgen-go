package cmd

import (
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode"

	"gabe565.com/pwgen/internal/config"
	"gabe565.com/pwgen/internal/wordlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	tmp, err := os.CreateTemp(t.TempDir(), "pwgen-test-config-*.toml")
	require.NoError(t, err)
	_ = tmp.Close()

	defaultArgs := []string{"--config=" + tmp.Name()}

	const defaultCount = 10

	type want struct {
		re        string
		lineCount int
		wordlist  wordlist.Meta
		split     string
		wordCount int
		nums      int
	}
	tests := []struct {
		name    string
		args    []string
		want    want
		wantErr require.ErrorAssertionFunc
	}{
		{
			"default",
			nil,
			want{re: `(\w+?\d?(-|$)){4,}`, lineCount: defaultCount, nums: 1},
			require.NoError,
		},
		{
			"count",
			[]string{"--count=20"},
			want{lineCount: 20, nums: 1},
			require.NoError,
		},
		{
			"template",
			[]string{"--template=abc"},
			want{re: `abc`, lineCount: defaultCount},
			require.NoError,
		},
		{
			"profile",
			[]string{"--profile=alpha"},
			want{re: `[A-Za-z]{32}`, lineCount: defaultCount},
			require.NoError,
		},
		{
			"profile param",
			[]string{"--profile=alpha:64"},
			want{re: `[A-Za-z]{64}`, lineCount: defaultCount},
			require.NoError,
		},
		{
			"default wordlist",
			[]string{"--profile=words:10", "--count=999"},
			want{split: " ", re: `[a-z\-]+`, wordCount: 10, lineCount: 999, wordlist: wordlist.Long},
			require.NoError,
		},
		{
			"short1 wordlist",
			[]string{"--profile=words:10", "--count=999", "--wordlist=short1"},
			want{split: " ", re: `[a-z\-]+`, wordCount: 10, lineCount: 999, wordlist: wordlist.Short1},
			require.NoError,
		},
		{
			"short2 wordlist",
			[]string{"--profile=words:10", "--count=999", "--wordlist=short2"},
			want{split: " ", re: `[a-z\-]+`, wordCount: 10, lineCount: 999, wordlist: wordlist.Short2},
			require.NoError,
		},
		{
			"invalid template",
			[]string{"--template={{ alpha"},
			want{lineCount: -1, nums: -1},
			require.Error,
		},
		{
			"invalid profile",
			[]string{"--profile=abc"},
			want{lineCount: -1, nums: -1},
			require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := New(WithContext(t.Context()))
			tt.args = append(defaultArgs, tt.args...)
			cmd.SetArgs(tt.args)
			var stdout strings.Builder
			cmd.SetOut(&stdout)

			tt.wantErr(t, cmd.Execute())

			wl, err := tt.want.wordlist.List()
			require.NoError(t, err)

			var re *regexp.Regexp
			if tt.want.re != "" {
				re, err = regexp.Compile("^" + tt.want.re + "$")
				require.NoError(t, err)
			}

			var lineCount int
			for _, line := range strings.Split(strings.TrimSpace(stdout.String()), "\n") {
				lineCount++

				if tt.want.nums != -1 {
					var nums int
					for _, r := range line {
						if unicode.IsDigit(r) {
							nums++
						}
					}
					assert.Equal(t, tt.want.nums, nums)
				}

				if tt.want.split == "" {
					if re != nil {
						assert.Regexp(t, re, line)
					}
				} else {
					assert.NotEmpty(t, wl)
					var wordCount int
					for _, word := range strings.Split(line, tt.want.split) {
						wordCount++
						if re != nil {
							assert.Regexp(t, re, word)
						}
						assert.True(t, slices.ContainsFunc(wl, func(s string) bool {
							return strings.EqualFold(s, word)
						}))
					}

					if tt.want.wordCount != -1 {
						assert.Equal(t, wordCount, tt.want.wordCount)
					}
				}
			}

			if tt.want.lineCount != -1 {
				assert.Equal(t, tt.want.lineCount, lineCount)
			}
		})
	}

	for _, profile := range config.New().Profiles.Named() {
		t.Run("profile "+profile.Name, func(t *testing.T) {
			cmd := New(WithContext(t.Context()))
			cmd.SetArgs(append(defaultArgs, "--profile="+profile.Name))
			var stdout strings.Builder
			cmd.SetOut(&stdout)
			require.NoError(t, cmd.Execute())
			require.NoError(t, err)
		})
	}
}
