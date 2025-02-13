package profiles

import (
	"cmp"
	"io"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"gabe565.com/pwgen/internal/config"
	"gabe565.com/pwgen/internal/funcmap"
	"gabe565.com/pwgen/internal/wordlist"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

const (
	name           = "profiles"
	formatKey      = "format"
	formatMarkdown = "markdown"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Default profile reference",

		ValidArgsFunction: cobra.NoFileCompletions,
	}
	cmd.SetHelpFunc(helpFunc)
	return cmd
}

func SetMarkdown(cmd *cobra.Command, v bool) {
	if cmd.Name() != name {
		var err error
		cmd, _, err = cmd.Find([]string{name})
		if err != nil {
			panic(err)
		}
	}

	if v {
		cmd.Annotations = map[string]string{formatKey: formatMarkdown}
	} else {
		delete(cmd.Annotations, formatKey)
	}
}

func helpFunc(cmd *cobra.Command, _ []string) {
	format := cmd.Annotations[formatKey]

	var result strings.Builder
	switch format {
	case formatMarkdown:
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
	t.AppendHeader(table.Row{"Name", "Example", "Template"})

	profiles := config.New().Profiles.Named()
	slices.SortStableFunc(profiles, func(a, b config.NamedProfile) int {
		return cmp.Compare(a.Name, b.Name)
	})
	words := wordlist.EFF_Long()
	tmpl := template.New("").Funcs(funcmap.New(words))

	for _, v := range profiles {
		name := v.Name
		if v.Param != 0 {
			name += ":" + strconv.Itoa(v.Param)
		}

		var buf strings.Builder
		tmpl, err := tmpl.New("").Parse(v.Template)
		if err == nil {
			_ = tmpl.Execute(&buf, v.Param)
		}

		switch format {
		case formatMarkdown:
			t.AppendRow(table.Row{"`" + name + "`", "<pre>" + buf.String() + "</pre>", "<pre>" + v.Template + "</pre>"})
		default:
			t.AppendRow(table.Row{name, buf.String(), v.Template})
		}
	}

	switch format {
	case formatMarkdown:
		result.WriteString(t.RenderMarkdown())
	default:
		result.WriteString(t.Render())
	}
	result.WriteByte('\n')
	_, _ = io.WriteString(cmd.OutOrStdout(), result.String())
}
