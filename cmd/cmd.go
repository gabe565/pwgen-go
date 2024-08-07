package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	pwgen_template "github.com/gabe565/pwgen-go/internal/template"
	"github.com/gabe565/pwgen-go/internal/wordlist"
	"github.com/spf13/cobra"
)

func New(opts ...Option) *cobra.Command {
	tmplSubcommand := NewProfiles(FormatText)

	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passphrases",
		Long: `Generate passphrases using the EFF Diceware Wordlists.
See https://www.eff.org/dice for details on the available wordlists.`,
		PreRunE: preRun,
		RunE:    run,

		ValidArgsFunction: cobra.NoFileCompletions,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(tmplSubcommand)

	cfg, _ := config.GetFilePretty()
	defaultCfg := config.NewDefault()

	cmd.Flags().String(config.FlagConfig, "", "Config file (default "+cfg+")")
	if err := cmd.RegisterFlagCompletionFunc(config.FlagConfig, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	}); err != nil {
		panic(err)
	}
	cmd.Flags().IntP(config.FlagCount, "c", defaultCfg.Count, "Number of passphrases to generate")
	if err := cmd.RegisterFlagCompletionFunc(config.FlagCount, cobra.NoFileCompletions); err != nil {
		panic(err)
	}

	cmd.Flags().String(config.FlagWordlist, config.WordlistLong, "Wordlist to use (one of: "+config.WordlistLong+", "+config.WordlistShort1+", "+config.WordlistShort2+")")
	if err := cmd.RegisterFlagCompletionFunc(config.FlagWordlist, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{config.WordlistLong, config.WordlistShort1, config.WordlistShort2}, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		panic(err)
	}

	cmd.Flags().StringP(config.FlagTemplate, "t", config.NewDefault().Template, `Template used to generate passphrases. If set, overrides the current profile.`)
	if err := cmd.RegisterFlagCompletionFunc(config.FlagTemplate, cobra.NoFileCompletions); err != nil {
		panic(err)
	}

	cmd.Flags().StringP(config.FlagProfile, "p", config.NewDefault().Template, `Generates passphrases using a preconfigured profile and an optional parameter. (see "pwgen profiles")`)
	if err := cmd.RegisterFlagCompletionFunc(config.FlagProfile, func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		conf, err := config.Load(cmd, false)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		named := make([]string, 0, len(conf.Profiles))
		wl, err := wordlist.New(conf.Wordlist)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		funcMap := pwgen_template.FuncMap(wl)
		var buf bytes.Buffer
		var longest int
		for name, v := range conf.Profiles {
			if v.Param != 0 {
				name += ":" + strconv.Itoa(v.Param)
			}
			if longest < len(name) {
				longest = len(name)
			}
		}
		for k, v := range conf.Profiles {
			if tmpl, err := template.New("").Funcs(funcMap).Parse(v.Template); err == nil {
				_ = tmpl.Execute(&buf, v.Param)
			}
			example := k
			if v.Param != 0 {
				example += ":" + strconv.Itoa(v.Param)
				if toComplete == k {
					k += ":"
				}
			}
			pad := strings.Repeat(" ", longest-len(example))
			named = append(named, fmt.Sprintf("%s\t%s%s -> %s", k, example, pad, buf.String()))
			buf.Reset()
		}
		return named, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}); err != nil {
		panic(err)
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

func preRun(cmd *cobra.Command, _ []string) error {
	conf, err := config.Load(cmd, true)
	if err != nil {
		return err
	}

	ctx := config.NewContext(cmd.Context(), conf)
	cmd.SetContext(ctx)

	return nil
}

var (
	ErrMissingConfig = errors.New("missing config")
	ErrInvalidFormat = errors.New("invalid format")
	ErrTemplate      = errors.New("template error")
)

func run(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		return ErrMissingConfig
	}

	wl, err := wordlist.New(conf.Wordlist)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Funcs(pwgen_template.FuncMap(wl)).Parse(conf.Template)
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
