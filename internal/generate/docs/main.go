package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gabe565.com/pwgen/cmd"
	"gabe565.com/utils/cobrax"
	"github.com/spf13/cobra/doc"
)

func main() {
	output := "./docs"

	if err := os.MkdirAll(output, 0o755); err != nil {
		panic(err)
	}

	root := cmd.New(cobrax.WithVersion("beta"), cmd.WithMarkdown())

	profCmd, _, err := root.Find([]string{"profiles"})
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := doc.GenMarkdown(root, &buf); err != nil {
		panic(err)
	}
	buf.WriteString("### SEE ALSO\n")
	buf.WriteString(
		fmt.Sprintf("* [%s %s](%s_%s.md)  - %s\n",
			root.Name(), "functions",
			root.Name(), "functions",
			"Template function reference",
		),
	)
	buf.WriteString(
		fmt.Sprintf("* [%s %s](%s_%s.md)  - %s\n",
			root.Name(), profCmd.Name(),
			root.Name(), profCmd.Name(),
			profCmd.Short,
		),
	)
	if err := os.WriteFile(filepath.Join(output, "pwgen.md"), buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	buf.Reset()

	var helpBuf strings.Builder
	profCmd.SetOut(&helpBuf)
	profCmd.HelpFunc()(profCmd, nil)

	buf.WriteString(
		fmt.Sprintf("# %s\n\n%s\n\n### SEE ALSO\n* [%s](%s.md)  - %s",
			profCmd.Name(),
			helpBuf.String(),
			root.Name(), root.Name(), root.Short,
		),
	)
	if err := os.WriteFile(filepath.Join(output, "pwgen_profiles.md"), buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
}
