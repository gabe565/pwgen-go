package cmd

import (
	"gabe565.com/pwgen/cmd/profiles"
	"gabe565.com/utils/cobrax"
	"github.com/spf13/cobra"
)

func WithRaw() cobrax.Option {
	return func(cmd *cobra.Command) {
		cmd.Long = long(true)
	}
}

func WithMarkdown() cobrax.Option {
	return func(cmd *cobra.Command) {
		WithRaw()(cmd)
		profiles.SetMarkdown(cmd, true)
	}
}
