package config

import (
	"strings"

	"github.com/gabe565/pwgen-go/internal/wordlist"
	"github.com/spf13/cobra"
)

const (
	FlagConfig = "config"

	FlagCount    = "count"
	FlagWordlist = "wordlist"
	FlagTemplate = "template"
	FlagProfile  = "profile"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	file, _ := GetFilePretty()
	fs.StringVar(&c.File, FlagConfig, c.File, `Config file (default "`+file+`")`)

	fs.IntP(FlagCount, "c", c.Count, "Number of passphrases to generate")
	fs.String(FlagWordlist, c.Wordlist.String(), "Wordlist to use (one of: "+strings.Join(wordlist.MetaStrings(), ", ")+")")
	cmd.Flags().StringP(FlagTemplate, "t", c.Template, `Template used to generate passphrases. If set, overrides the current profile.`)
	fs.StringP(FlagProfile, "p", c.Template, `Generates passphrases using a preconfigured profile and an optional parameter. (see "pwgen profiles")`)
}
