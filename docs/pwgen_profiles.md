# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Example | Template |
| --- | --- | --- |
| `alpha:32` | <code>ogopMCZXGJiXwGUomaDRChKBRkmqxXej</code> | `{{ alpha . }}` |
| `alphanum:32` | <code>V0cwh6NuqNh3dQwYtppteI8aypyTXPjm</code> | `{{ alphaNum . }}` |
| `ascii:32` | <code>+6nTemiq@d(X@/;5*G[q0U;3:>SY~\|<E</code> | `{{ ascii . }}` |
| `diceware:4` | <code>Shorter-Scaling-Studied-Uncured4</code> | `{{ wordsWithNum . \| join "-" \| title }}` |
| `django:0` | <code>$pfy@*uy_$21ebp)9y!c%%=j=75w_d8n0&52g9-vv@=*9l@hu(</code> | `{{ randFromStr "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)" 50 }}` |
| `laravel:0` | <code>base64:p19Unv0Xmy971x6bw8FRZ7sjYJOqlN5FU3MHO50h6/4=</code> | `base64:{{ binary 32 \| b64enc }}` |
| `pin:6` | <code>472054</code> | `{{ num . }}` |
| `words:4` | <code>caress hamstring dubbed hut</code> | `{{ words . \| join " " }}` |


### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases