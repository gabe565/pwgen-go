# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Template | Example |
| --- | --- | --- |
| `alpha:32` | `{{ alpha . }}` | <code>xtkMnIVuuqIOLOSaGWCxqBVuTgVSbnto</code> |
| `alphanum:32` | `{{ alphaNum . }}` | <code>2y6WdoPVTjC0YP6OGQtd8ER46FX5zhxe</code> |
| `ascii:32` | `{{ ascii . }}` | <code>}$0e3\|P3HRm8e=\wm$,&-9B)~e<ebw/E</code> |
| `diceware:4` | `{{ wordsWithNum . \| join "-" \| title }}` | <code>Fever-Barge-Excretory-Asleep5</code> |
| `laravel:0` | `base64:{{ binary 32 \| b64enc }}` | <code>base64:3fKizKd3cqv9MwzJ+JEII5Ir/872NziLb5qZHy+UGlY=</code> |
| `pin:6` | `{{ num . }}` | <code>492042</code> |
| `words:4` | `{{ words . \| join " " }}` | <code>freewill outpour scuff dwindle</code> |

### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases