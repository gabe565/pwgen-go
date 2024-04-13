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
      --config string     Config file (default $HOME/.config/pwgen-go/config.toml)
  -c, --count int         Number of passphrases to generate (default 10)
  -h, --help              help for pwgen
  -t, --template string   Template used to generate passphrases. Either a Go template or a named template (see "pwgen templates"). (default "diceware-3")
      --wordlist string   Wordlist to use (one of: long, short1, short2) (default "long")
```

