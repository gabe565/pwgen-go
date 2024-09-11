package cmd

import (
	"github.com/gabe565/pwgen-go/cmd/profiles"
	"github.com/spf13/cobra"
)

type Option func(cmd *cobra.Command)

func WithVersion(version string) Option {
	return func(cmd *cobra.Command) {
		cmd.Version = buildVersion(version)
		cmd.InitDefaultVersionFlag()
	}
}

func WithRaw() Option {
	return func(cmd *cobra.Command) {
		cmd.Long = long(true)
	}
}

func WithMarkdown() Option {
	return func(cmd *cobra.Command) {
		WithRaw()(cmd)
		profiles.SetMarkdown(cmd, true)
	}
}
