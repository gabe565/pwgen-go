package cmd

import (
	"github.com/gabe565/pwgen-go/cmd/template"
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/spf13/cobra"
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

	cmd.AddCommand(
		template.New(),
	)
	return cmd
}
