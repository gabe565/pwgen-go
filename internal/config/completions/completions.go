package completions

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/gabe565/pwgen-go/internal/funcmap"
	"github.com/gabe565/pwgen-go/internal/wordlist"
	"github.com/spf13/cobra"
)

func Register(cmd *cobra.Command) {
	if err := errors.Join(
		cmd.RegisterFlagCompletionFunc(config.FlagConfig,
			func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
				return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
			},
		),
		cmd.RegisterFlagCompletionFunc(config.FlagCount, cobra.NoFileCompletions),
		cmd.RegisterFlagCompletionFunc(config.FlagWordlist,
			func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
				lists := wordlist.MetaValues()
				values := make([]string, 0, len(lists))
				for _, wl := range lists {
					values = append(values, wl.String()+"\t"+strings.ReplaceAll(wl.Description(), "\n", " "))
				}
				return values, cobra.ShellCompDirectiveNoFileComp
			},
		),
		cmd.RegisterFlagCompletionFunc(config.FlagTemplate, cobra.NoFileCompletions),
		cmd.RegisterFlagCompletionFunc(config.FlagProfile, completeProfile),
	); err != nil {
		panic(err)
	}
}

func completeProfile(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	conf, err := config.Load(cmd, false)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	named := make([]string, 0, len(conf.Profiles))
	wl, err := conf.Wordlist.List()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	funcMap := funcmap.New(wl)
	var buf bytes.Buffer
	var longest int
	for name, v := range conf.Profiles {
		if v.Param != 0 {
			name += ":" + strconv.Itoa(v.Param)
		}
		if longest < len(name) {
			longest = len(name)
		}
	}
	for k, v := range conf.Profiles {
		if tmpl, err := template.New("").Funcs(funcMap).Parse(v.Template); err == nil {
			_ = tmpl.Execute(&buf, v.Param)
		}
		example := k
		if v.Param != 0 {
			example += ":" + strconv.Itoa(v.Param)
			if toComplete == k {
				k += ":"
			}
		}
		pad := strings.Repeat(" ", longest-len(example))
		named = append(named, fmt.Sprintf("%s\t%s%s -> %s", k, example, pad, buf.String()))
		buf.Reset()
	}
	return named, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
}
