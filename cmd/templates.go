package cmd

import (
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type Format uint8

const (
	FormatText Format = iota
	FormatMarkdown
)

func NewTemplates(format Format) *cobra.Command {
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
		switch format {
		case FormatText:
			t.AppendRow(table.Row{k, v})
		case FormatMarkdown:
			t.AppendRow(table.Row{"`" + k + "`", "`" + v + "`"})
		}
	}
	t.SortBy([]table.SortBy{{Number: 1}})

	switch format {
	case FormatText:
		return &cobra.Command{
			Use:   "templates",
			Short: "Default named template reference",
			Long: "The --template flag can be a raw Go template, or it can be a named template.\n\n" +
				"Default Named Templates:\n" + t.Render(),
			ValidArgsFunction: cobra.NoFileCompletions,
		}
	case FormatMarkdown:
		return &cobra.Command{
			Use:   "templates",
			Short: "Default named template reference",
			Long: "The `--template` flag can be a raw Go template, or it can be a named template.\n\n" +
				"## Default Named Templates\n\n" + t.RenderMarkdown(),
			ValidArgsFunction: cobra.NoFileCompletions,
		}
	default:
		panic("invalid format")
	}
}
