package cmd

import (
	"strconv"

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

func NewProfiles(format Format) *cobra.Command {
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
	for k, v := range conf.Profiles {
		name := k + ":" + strconv.Itoa(v.Param)
		switch format {
		case FormatText:
			t.AppendRow(table.Row{name, v.Template})
		case FormatMarkdown:
			t.AppendRow(table.Row{"`" + name + "`", "`" + v.Template + "`"})
		}
	}
	t.SortBy([]table.SortBy{{Number: 1}})

	switch format {
	case FormatText:
		return &cobra.Command{
			Use:   "profiles",
			Short: "Default profile reference",
			Long: "The --profile flag lets you use preconfigured templates with an optional colon-separated parameter.\n\n" +
				"Default Profiles:\n" + t.Render(),
			ValidArgsFunction: cobra.NoFileCompletions,
		}
	case FormatMarkdown:
		return &cobra.Command{
			Use:   "profiles",
			Short: "Default profile reference",
			Long: "The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.\n\n" +
				"## Default Profiles\n\n" + t.RenderMarkdown(),
			ValidArgsFunction: cobra.NoFileCompletions,
		}
	default:
		panic("invalid format")
	}
}
