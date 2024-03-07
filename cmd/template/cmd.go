package template

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/gabe565/pwgen-go/internal/util"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "template",
		Short:   "Generates passwords from a template",
		PreRunE: preRun,
		RunE:    run,
	}

	cmd.Flags().StringP("template", "t", config.NewDefault().Template, "Go template that generates a password")

	return cmd
}

func preRun(cmd *cobra.Command, args []string) error {
	conf, err := config.Load(cmd)
	if err != nil {
		return err
	}

	ctx := config.NewContext(cmd.Context(), conf)
	cmd.SetContext(ctx)

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		return fmt.Errorf("missing config")
	}

	tmpl, err := template.New("").Funcs(util.TemplateFuncMap()).Parse(conf.Template)
	if err != nil {
		return fmt.Errorf("invalid format: %w", err)
	}

	var buf strings.Builder
	for i := 0; i < conf.Count; i += 1 {
		if err := tmpl.Execute(&buf, nil); err != nil {
			return fmt.Errorf("template error: %w", err)
		}
		fmt.Println(buf.String())
		buf.Reset()
	}

	return nil
}
