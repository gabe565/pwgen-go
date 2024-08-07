package cmd

import (
	"github.com/gabe565/pwgen-go/cmd/profiles"
	"github.com/spf13/cobra"
)

type Option func(cmd *cobra.Command)

func WithVersion(version string) Option {
	return func(cmd *cobra.Command) {
		cmd.Version = buildVersion(version)
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

		profileCmd, _, err := cmd.Find([]string{"profiles"})
		if err != nil {
			panic(err)
		}
		profileCmd.Annotations = map[string]string{"format": profiles.FormatMarkdown}
	}
}
