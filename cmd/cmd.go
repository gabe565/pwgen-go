package cmd

import (
	"os"

	"github.com/gabe565/pwgen-go/cmd/template"
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passwords",

		DisableAutoGenTag: true,
	}

	cfg, _ := config.GetFilePretty()
	cmd.PersistentFlags().String("config", "", "Config file (default "+cfg+")")
	cmd.PersistentFlags().IntP("count", "c", config.NewDefault().Count, "Number of passwords to generate")

	template := template.New()
	cmd.AddCommand(
		template,
	)

	// default cmd if no cmd is given
	subCmd, _, err := cmd.Find(os.Args[1:])
	if err == nil && subCmd.Use == cmd.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp {
		args := append([]string{template.Use}, os.Args[1:]...)
		cmd.SetArgs(args)
	}

	return cmd
}
