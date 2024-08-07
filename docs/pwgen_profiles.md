# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Template | Example |
| --- | --- | --- |
| `alpha:32` | `{{ alpha . }}` | <code>ogopMCZXGJiXwGUomaDRChKBRkmqxXej</code> |
| `alphanum:32` | `{{ alphaNum . }}` | <code>V0cwh6NuqNh3dQwYtppteI8aypyTXPjm</code> |
| `ascii:32` | `{{ ascii . }}` | <code>+6nTemiq@d(X@/;5*G[q0U;3:>SY~\|<E</code> |
| `diceware:4` | `{{ wordsWithNum . \| join "-" \| title }}` | <code>Shorter-Scaling-Studied-Uncured4</code> |
| `laravel:0` | `base64:{{ binary 32 \| b64enc }}` | <code>base64:WFiyK3glryeARDrWbLj/YYDCI491HUEFluAWR0rEFCM=</code> |
| `pin:6` | `{{ num . }}` | <code>586517</code> |
| `words:4` | `{{ words . \| join " " }}` | <code>uneatable hungry imply shininess</code> |


### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases