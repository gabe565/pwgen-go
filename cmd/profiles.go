package cmd

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	pwgen_template "github.com/gabe565/pwgen-go/internal/template"
	"github.com/gabe565/pwgen-go/internal/wordlist"
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

	profiles := conf.Profiles.Named()
	slices.SortStableFunc(profiles, func(a, b config.NamedProfile) int {
		return cmp.Compare(a.Name, b.Name)
	})

	words := wordlist.EFF_Long()

	t.AppendHeader(table.Row{"Name", "Template", "Example"})
	for _, v := range profiles {
		name := v.Name + ":" + strconv.Itoa(v.Param)

		var buf strings.Builder
		tmpl, err := template.New("").Funcs(pwgen_template.FuncMap(words)).Parse(v.Template)
		if err == nil {
			_ = tmpl.Execute(&buf, v.Param)
		}

		switch format {
		case FormatText:
			t.AppendRow(table.Row{name, v.Template, buf.String()})
		case FormatMarkdown:
			t.AppendRow(table.Row{"`" + name + "`", "`" + v.Template + "`", "<code>" + buf.String() + "</code>"})
		}
	}

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
