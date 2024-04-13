package cmd

import (
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func NewTemplates() *cobra.Command {
	conf := config.NewDefault()

	t := table.NewWriter()
	style := table.StyleLight
	style.Box.Left = "  " + style.Box.Left
	style.Box.LeftSeparator = "  " + style.Box.LeftSeparator
	style.Box.BottomLeft = "  " + style.Box.BottomLeft
	style.Box.TopLeft = "  " + style.Box.TopLeft
	style.Color.Border = text.Colors{text.FgHiBlack}
	style.Color.Separator = style.Color.Border
	t.SetStyle(style)

	t.AppendHeader(table.Row{"Name", "Template"})
	for k, v := range conf.Templates {
		t.AppendRow(table.Row{k, v})
	}
	t.SortBy([]table.SortBy{{Number: 1}})

	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Default named template reference",
		Long: `The --template flag can be a raw Go template, or it can be a named template.

Default Named Templates:
` + t.Render(),

		ValidArgsFunction: cobra.NoFileCompletions,
	}
	return cmd
}
