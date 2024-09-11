package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"text/template"

	"github.com/gabe565/pwgen-go/cmd/profiles"
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/gabe565/pwgen-go/internal/config/completions"
	"github.com/gabe565/pwgen-go/internal/funcmap"
	"github.com/gabe565/pwgen-go/internal/wordlist"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

func New(opts ...Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passphrases",
		Long:  long(false),
		RunE:  run,

		ValidArgsFunction: cobra.NoFileCompletions,
		DisableAutoGenTag: true,
	}

	conf := config.New()
	conf.RegisterFlags(cmd)
	cmd.SetContext(config.NewContext(context.Background(), conf))
	completions.Register(cmd)

	cmd.AddCommand(profiles.New())

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

func long(rawText bool) string {
	var link string
	if rawText {
		link = "EFF Diceware Wordlists"
	} else {
		link = termenv.Hyperlink("https://www.eff.org/dice", "EFF Diceware Wordlists")
	}
	return "Generate passphrases using the " + link + "."
}

var (
	ErrInvalidFormat = errors.New("invalid format")
	ErrTemplate      = errors.New("template error")
)

func run(cmd *cobra.Command, _ []string) error {
	conf, err := config.Load(cmd, true)
	if err != nil {
		return err
	}

	cmd.SilenceUsage = true

	wl, err := wordlist.New(conf.Wordlist)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Funcs(funcmap.New(wl)).Parse(conf.Template)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidFormat, err)
	}

	w := bufio.NewWriter(cmd.OutOrStdout())
	for range conf.Count {
		if err := tmpl.Execute(w, conf.Param); err != nil {
			return fmt.Errorf("%w: %w", ErrTemplate, err)
		}
	}
	return w.Flush()
}
