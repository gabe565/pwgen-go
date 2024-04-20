# templates

The `--template` flag can be a raw Go template, or it can be a named template.

## Default Named Templates

| Name | Template |
| --- | --- |
| `alpha-16` | `{{ randAlpha 16 }}` |
| `alpha-32` | `{{ randAlpha 32 }}` |
| `alpha-64` | `{{ randAlpha 64 }}` |
| `ascii-16` | `{{ randAscii 16 }}` |
| `ascii-32` | `{{ randAscii 32 }}` |
| `ascii-64` | `{{ randAscii 64 }}` |
| `diceware-3` | `{{ wordsWithNumber 3 \| join "-" \| title }}` |
| `diceware-5` | `{{ wordsWithNumber 5 \| join "-" \| title }}` |
| `diceware-6` | `{{ wordsWithNumber 6 \| join "-" \| title }}` |
| `pin-4` | `{{ num 4 }}` |
| `pin-6` | `{{ num 6 }}` |

### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases