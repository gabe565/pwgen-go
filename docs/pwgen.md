## pwgen

Generate passphrases

### Synopsis

Generate passphrases using the EFF Diceware Wordlists.

```
pwgen [flags]
```

### Options

```
      --config string     Config file (default $HOME/.config/pwgen-go/config.toml)
  -c, --count int         Number of passphrases to generate (default 10)
  -h, --help              help for pwgen
  -p, --profile string    Generates passphrases using a preconfigured profile and an optional parameter. (see "pwgen profiles")
  -t, --template string   Template used to generate passphrases. If set, overrides the current profile.
  -v, --version           version for pwgen
      --wordlist string   Wordlist to use (one of: long, short1, short2) (default "long")
```

### SEE ALSO
* [pwgen functions](pwgen_functions.md)  - Template function reference
* [pwgen profiles](pwgen_profiles.md)  - Default profile reference
