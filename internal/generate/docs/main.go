package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gabe565/pwgen-go/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	output := "./docs"

	if err := os.MkdirAll(output, 0o755); err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	root := cmd.New("", "")
	tmpl := cmd.NewProfiles(cmd.FormatMarkdown)
	if err := doc.GenMarkdown(root, &buf); err != nil {
		panic(err)
	}
	buf.WriteString("### SEE ALSO\n")
	buf.WriteString(fmt.Sprintf("* [%s %s](%s_%s.md)  - %s\n", root.Name(), "functions", root.Name(), "functions", "Template function reference"))
	buf.WriteString(fmt.Sprintf("* [%s %s](%s_%s.md)  - %s\n", root.Name(), tmpl.Name(), root.Name(), tmpl.Name(), tmpl.Short))
	if err := os.WriteFile(filepath.Join(output, "pwgen.md"), buf.Bytes(), 0o644); err != nil {
		panic(err)
	}

	buf.Reset()
	buf.WriteString(fmt.Sprintf("# %s\n\n%s\n\n### SEE ALSO\n* [%s](%s.md)  - %s", tmpl.Name(), tmpl.Long, root.Name(), root.Name(), root.Short))
	if err := os.WriteFile(filepath.Join(output, "pwgen_profiles.md"), buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
}
