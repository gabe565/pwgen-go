package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	pwgen_template "github.com/gabe565/pwgen-go/internal/template"
	"github.com/spf13/cobra"
)

func New(version, commit string) *cobra.Command {
	tmplSubcommand := NewTemplates()

	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passphrases",
		Long: `Generate passphrases using the EFF Diceware Wordlists.
See https://www.eff.org/dice for details on the available wordlists.`,
		Version: buildVersion(version, commit),
		PreRunE: preRun,
		RunE:    run,

		ValidArgsFunction: cobra.NoFileCompletions,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(tmplSubcommand)

	cfg, _ := config.GetFilePretty()
	defaultCfg := config.NewDefault()

	cmd.Flags().String("config", "", "Config file (default "+cfg+")")
	if err := cmd.RegisterFlagCompletionFunc("config", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	}); err != nil {
		panic(err)
	}
	cmd.Flags().IntP("count", "c", defaultCfg.Count, "Number of passphrases to generate")
	if err := cmd.RegisterFlagCompletionFunc("count", cobra.NoFileCompletions); err != nil {
		panic(err)
	}

	cmd.Flags().String("wordlist", "long", "Wordlist to use (one of: long, short1, short2)")
	if err := cmd.RegisterFlagCompletionFunc("wordlist", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{config.WordlistLong, config.WordlistShort1, config.WordlistShort2}, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		panic(err)
	}

	cmd.Flags().StringP("template", "t", config.NewDefault().Template, `Template used to generate passphrases. Either a Go template or a named template (see "pwgen templates").`)
	if err := cmd.RegisterFlagCompletionFunc("template", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		conf, err := config.Load(cmd, false)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		named := make([]string, 0, len(conf.Templates))
		funcMap := pwgen_template.FuncMap(conf)
		var buf bytes.Buffer
		for k, v := range conf.Templates {
			if tmpl, err := template.New("").Funcs(funcMap).Parse(v); err == nil {
				_ = tmpl.Execute(&buf, nil)
			}
			named = append(named, k+"\texample: "+buf.String())
			buf.Reset()
		}
		return named, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		panic(err)
	}

	return cmd
}

func buildVersion(version, commit string) string {
	if commit != "" {
		version += " (" + commit + ")"
	}
	return version
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
)

func run(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		return ErrMissingConfig
	}

	tmpl, err := template.New("").Funcs(pwgen_template.FuncMap(conf)).Parse(conf.Template)
	if err != nil {
		return ErrInvalidFormat
	}

	var buf bytes.Buffer
	for range conf.Count {
		if err := tmpl.Execute(&buf, nil); err != nil {
			return fmt.Errorf("template error: %w", err)
		}
		_, _ = io.Copy(cmd.OutOrStdout(), &buf)
	}

	return nil
}
