## pwgen template

Generates passwords from a template (default)

```
pwgen template [flags]
```

### Options

```
  -h, --help              help for template
  -t, --template string   Go template that generates a password (default "{{ words 3 | join \"-\" | title }}{{ number 1 }}")
```

### Options inherited from parent commands

```
      --config string   Config file (default $HOME/.config/pwgen-go/config.toml)
  -c, --count int       Number of passwords to generate (default 10)
```

### SEE ALSO

* [pwgen](pwgen.md)	 - Generate passwords

