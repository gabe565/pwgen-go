package cmd

import (
	"os"

	"github.com/gabe565/pwgen-go/cmd/template"
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func New(version, commit string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passphrases",
		Long: `Generate passphrases using the EFF Diceware Wordlists.
See https://www.eff.org/dice for details on the available wordlists.`,
		Version: buildVersion(version, commit),

		DisableAutoGenTag: true,
	}

	cfg, _ := config.GetFilePretty()
	defaultCfg := config.NewDefault()
	cmd.PersistentFlags().String("config", "", "Config file (default "+cfg+")")
	_ = cmd.RegisterFlagCompletionFunc("config", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	})
	cmd.PersistentFlags().IntP("count", "c", defaultCfg.Count, "Number of passphrases to generate")
	cmd.PersistentFlags().String("wordlist", "long", "Wordlist to use (one of: long, short1, short2)")
	_ = cmd.RegisterFlagCompletionFunc("wordlist", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{config.WordlistLong, config.WordlistShort1, config.WordlistShort2}, cobra.ShellCompDirectiveNoFileComp
	})

	template := template.New()
	cmd.AddCommand(
		template,
	)

	// default cmd if no cmd is given
	cmd.InitDefaultVersionFlag()
	subCmd, _, err := cmd.Find(os.Args[1:])
	if err == nil && subCmd.Use == cmd.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp {
		if versionFlag, err := cmd.Flags().GetBool("version"); err == nil && !versionFlag {
			args := append([]string{template.Use}, os.Args[1:]...)
			cmd.SetArgs(args)
		}
	}

	return cmd
}

func buildVersion(version, commit string) string {
	if commit != "" {
		version += " (" + commit + ")"
	}
	return version
}
