package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"text/template"

	"gabe565.com/pwgen/cmd/profiles"
	"gabe565.com/pwgen/internal/config"
	"gabe565.com/pwgen/internal/config/completions"
	"gabe565.com/pwgen/internal/funcmap"
	"gabe565.com/pwgen/internal/wordlist"
	"gabe565.com/utils/cobrax"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwgen [profile | template]",
		Short: "Generate passphrases",
		Long:  long(false),
		RunE:  run,
		Args:  cobra.MaximumNArgs(1),

		ValidArgsFunction: validArgs,
		DisableAutoGenTag: true,
		SilenceErrors:     true,
		SilenceUsage:      true,
	}

	conf := config.New()
	conf.RegisterFlags(cmd)
	completions.Register(cmd)

	cmd.AddCommand(profiles.New())

	for _, opt := range opts {
		opt(cmd)
	}
	if cmd.Context() == nil {
		cmd.SetContext(context.Background())
	}
	cmd.SetContext(config.NewContext(cmd.Context(), conf))

	return cmd
}

func validArgs(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return completions.Profile(cmd, args, toComplete)
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

func run(cmd *cobra.Command, args []string) error {
	conf, err := config.Load(cmd, args, true)
	if err != nil {
		return err
	}

	wl, err := wordlist.Load(conf.Wordlist)
	if err != nil {
		return err
	}

	tmpl := template.New("").Funcs(funcmap.New(wl))
	for name, profile := range conf.Profiles {
		if _, err := tmpl.New(name).Parse(profile.Template); err != nil {
			return err
		}
	}

	tmpl, err = tmpl.New("").Parse(conf.Template)
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
