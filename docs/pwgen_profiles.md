# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Template | Example |
| --- | --- | --- |
| `alpha:32` | `{{ alpha . }}` | <code>ogopMCZXGJiXwGUomaDRChKBRkmqxXej</code> |
| `alphanum:32` | `{{ alphaNum . }}` | <code>V0cwh6NuqNh3dQwYtppteI8aypyTXPjm</code> |
| `ascii:32` | `{{ ascii . }}` | <code>+6nTemiq@d(X@/;5*G[q0U;3:>SY~\|<E</code> |
| `diceware:4` | `{{ wordsWithNum . \| join "-" \| title }}` | <code>Shorter-Scaling-Studied-Uncured4</code> |
| `django:0` | `{{ randFromStr "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)" 50 }}` | <code>$pfy@*uy_$21ebp)9y!c%%=j=75w_d8n0&52g9-vv@=*9l@hu(</code> |
| `laravel:0` | `base64:{{ binary 32 \| b64enc }}` | <code>base64:p19Unv0Xmy971x6bw8FRZ7sjYJOqlN5FU3MHO50h6/4=</code> |
| `pin:6` | `{{ num . }}` | <code>472054</code> |
| `words:4` | `{{ words . \| join " " }}` | <code>caress hamstring dubbed hut</code> |


### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases