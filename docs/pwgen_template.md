## pwgen template

Generates passphrases from a template (default)

```
pwgen template [flags]
```

### Options

```
  -h, --help              help for template
  -t, --template string   Go template that generates a passphrase (default "{{ wordsWithNumber 3 | join \"-\" | title }}")
```

### Options inherited from parent commands

```
      --config string   Config file (default $HOME/.config/pwgen-go/config.toml)
  -c, --count int       Number of passphrases to generate (default 10)
```

### SEE ALSO

* [pwgen](pwgen.md)	 - Generate passphrases

