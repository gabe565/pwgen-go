# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Template |
| --- | --- |
| `alpha:32` | `{{ randAlpha . }}` |
| `ascii:32` | `{{ randAscii . }}` |
| `diceware:4` | `{{ wordsWithNumber . \| join "-" \| title }}` |
| `pin:6` | `{{ num . }}` |
| `words:4` | `{{ words . \| join " " }}` |

### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases