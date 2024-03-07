package template

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/util"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Generates passwords from a template",
		RunE:  run,
	}

	cmd.Flags().StringP("template", "t", `{{ randWords 3 | join "-" | title }}{{ randNumeric 1 }}`, "Go template that generates a password")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	tmplSrc, err := cmd.Flags().GetString("template")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("").Funcs(util.TemplateFuncMap()).Parse(tmplSrc)
	if err != nil {
		return fmt.Errorf("invalid format: %w", err)
	}

	var buf strings.Builder

	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i += 1 {
		if err := tmpl.Execute(&buf, nil); err != nil {
			return fmt.Errorf("template error: %w", err)
		}
		fmt.Println(buf.String())
		buf.Reset()
	}

	return nil
}
