# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Template |
| --- | --- |
| `alpha:32` | `{{ alpha . }}` |
| `alphanum:32` | `{{ alphaNum . }}` |
| `ascii:32` | `{{ ascii . }}` |
| `diceware:4` | `{{ wordsWithNum . \| join "-" \| title }}` |
| `laravel:0` | `base64:{{ binary 32 \| b64enc }}` |
| `pin:6` | `{{ num . }}` |
| `words:4` | `{{ words . \| join " " }}` |

### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases