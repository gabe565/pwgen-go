# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Example | Template |
| --- | --- | --- |
| `alpha:32` | <pre>ogopMCZXGJiXwGUomaDRChKBRkmqxXej</pre> | <pre>{{ alpha . }}</pre> |
| `alphanum:32` | <pre>V0cwh6NuqNh3dQwYtppteI8aypyTXPjm</pre> | <pre>{{ alphaNum . }}</pre> |
| `ascii:32` | <pre>+6nTemiq@d(X@/;5*G[q0U;3:>SY~\|<E</pre> | <pre>{{ ascii . }}</pre> |
| `diceware:4` | <pre>Shorter-Scaling-Studied-Uncured4</pre> | <pre>{{ wordsWithNum . \| join "-" \| title }}</pre> |
| `django` | <pre>$pfy@*uy_$21ebp)9y!c%%=j=75w_d8n0&52g9-vv@=*9l@hu(</pre> | <pre>{{ randFromStr "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)" 50 }}</pre> |
| `laravel` | <pre>base64:p19Unv0Xmy971x6bw8FRZ7sjYJOqlN5FU3MHO50h6/4=</pre> | <pre>base64:{{ binary 32 \| b64enc }}</pre> |
| `pin:6` | <pre>472054</pre> | <pre>{{ num . }}</pre> |
| `words:4` | <pre>caress hamstring dubbed hut</pre> | <pre>{{ words . \| join " " }}</pre> |


### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases