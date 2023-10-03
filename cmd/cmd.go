package cmd

import (
	"github.com/gabe565/pwgen-go/cmd/template"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate passwords",
	}

	cmd.PersistentFlags().IntP("count", "c", 10, "Number of passwords to generate")

	cmd.AddCommand(
		template.New(),
	)
	return cmd
}
