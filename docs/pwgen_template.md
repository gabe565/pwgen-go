## pwgen template

Generates passwords from a template

```
pwgen template [flags]
```

### Options

```
  -h, --help              help for template
  -t, --template string   Go template that generates a password (default "{{ randWords 3 | join \"-\" | title }}{{ randNumeric 1 }}")
```

### Options inherited from parent commands

```
  -c, --count int   Number of passwords to generate (default 10)
```

### SEE ALSO

* [pwgen](pwgen.md)	 - Generate passwords

