package profiles

import (
	"cmp"
	"io"
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

const FormatMarkdown = "markdown"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Default profile reference",

		ValidArgsFunction: cobra.NoFileCompletions,
	}
	cmd.SetHelpFunc(helpFunc)
	return cmd
}

func helpFunc(cmd *cobra.Command, _ []string) {
	format := cmd.Annotations["format"]

	var result strings.Builder
	switch format {
	case FormatMarkdown:
		result.WriteString("The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.\n\n" +
			"## Default Profiles\n\n")
	default:
		result.WriteString("The --profile flag lets you use preconfigured templates with an optional colon-separated parameter.\n\n" +
			"Default Profiles:\n")
	}

	t := table.NewWriter()
	style := table.StyleLight
	style.Box.Left = "  " + style.Box.Left
	style.Box.LeftSeparator = "  " + style.Box.LeftSeparator
	style.Box.BottomLeft = "  " + style.Box.BottomLeft
	style.Box.TopLeft = "  " + style.Box.TopLeft
	style.Color.Border = text.Colors{text.FgHiBlack}
	style.Color.Separator = style.Color.Border
	t.SetStyle(style)

	profiles := config.NewDefault().Profiles.Named()
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
		case FormatMarkdown:
			t.AppendRow(table.Row{"`" + name + "`", "`" + v.Template + "`", "<code>" + buf.String() + "</code>"})
		default:
			t.AppendRow(table.Row{name, v.Template, buf.String()})
		}
	}

	switch format {
	case FormatMarkdown:
		result.WriteString(t.RenderMarkdown())
	default:
		result.WriteString(t.Render())
	}
	result.WriteByte('\n')
	_, _ = io.WriteString(cmd.OutOrStdout(), result.String())
}
