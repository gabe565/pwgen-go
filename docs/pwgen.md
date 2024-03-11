## pwgen

Generate passphrases

### Synopsis

Generate passphrases using the EFF Diceware Wordlists.
See https://www.eff.org/dice for details on the available wordlists.

```
pwgen [flags]
```

### Options

```
      --completion string   Output command-line completion code for the specified shell. Can be 'bash', 'zsh', 'fish', or 'powershell'.
      --config string       Config file (default $HOME/.config/pwgen-go/config.toml)
  -c, --count int           Number of passphrases to generate (default 10)
  -h, --help                help for pwgen
  -t, --template string     Go template that generates a password (default "{{ wordsWithNumber 3 | join \"-\" | title }}")
      --wordlist string     Wordlist to use (one of: long, short1, short2) (default "long")
```

