package cmd

import (
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	pwgen_template "github.com/gabe565/pwgen-go/internal/template"
	"github.com/spf13/cobra"
)

func New(version, commit string) *cobra.Command {
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

	cfg, _ := config.GetFilePretty()
	defaultCfg := config.NewDefault()
	registerCompletionFlag(cmd)
	cmd.PersistentFlags().String("config", "", "Config file (default "+cfg+")")
	if err := cmd.RegisterFlagCompletionFunc("config", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	}); err != nil {
		panic(err)
	}
	cmd.PersistentFlags().IntP("count", "c", defaultCfg.Count, "Number of passphrases to generate")
	cmd.PersistentFlags().String("wordlist", "long", "Wordlist to use (one of: long, short1, short2)")
	if err := cmd.RegisterFlagCompletionFunc("wordlist", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{config.WordlistLong, config.WordlistShort1, config.WordlistShort2}, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		panic(err)
	}
	cmd.Flags().StringP("template", "t", config.NewDefault().Template, "Go template that generates a password")

	return cmd
}

func buildVersion(version, commit string) string {
	if commit != "" {
		version += " (" + commit + ")"
	}
	return version
}

func preRun(cmd *cobra.Command, _ []string) error {
	completionFlag, err := cmd.Flags().GetString(CompletionFlag)
	if err != nil {
		panic(err)
	}
	if completionFlag != "" {
		return nil
	}

	conf, err := config.Load(cmd)
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

func run(cmd *cobra.Command, args []string) error {
	completionFlag, err := cmd.Flags().GetString(CompletionFlag)
	if err != nil {
		panic(err)
	}
	if completionFlag != "" {
		return completion(cmd, args)
	}

	cmd.SilenceUsage = true

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		return ErrMissingConfig
	}

	tmpl, err := template.New("").Funcs(pwgen_template.FuncMap(conf)).Parse(conf.Template)
	if err != nil {
		return ErrInvalidFormat
	}

	var buf strings.Builder
	for range conf.Count {
		if err := tmpl.Execute(&buf, nil); err != nil {
			return fmt.Errorf("template error: %w", err)
		}
		//nolint:forbidigo
		fmt.Println(buf.String())
		buf.Reset()
	}

	return nil
}
