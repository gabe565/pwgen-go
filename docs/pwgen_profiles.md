# profiles

The `--profile` flag lets you use preconfigured templates with an optional colon-separated parameter.

## Default Profiles

| Name | Example | Template |
| --- | --- | --- |
| `alpha:32` | <pre>ogopMCZXGJiXwGUomaDRChKBRkmqxXej</pre> | <pre>{{ alpha . }}</pre> |
| `alphanum:32` | <pre>V0cwh6NuqNh3dQwYtppteI8aypyTXPjm</pre> | <pre>{{ alphaNum . }}</pre> |
| `ascii:32` | <pre>+6nTemiq@d(X@/;5*G[q0U;3:>SY~\|<E</pre> | <pre>{{ ascii . }}</pre> |
| `base64:32` | <pre>HQ0mLMdPKmhYWLIreCWvJ4BEOtZsuP9hgMIjj3UdQQU=</pre> | <pre>{{ binary . \| b64enc }}</pre> |
| `diceware:4` | <pre>Kept9-Reforest-Profanity-Uptight</pre> | <pre>{{ wordsWithNum . \| join "-" \| title }}</pre> |
| `django` | <pre>0&52g9-vv@=*9l@hu((%iolnrzgwb=^bu(t5ur0@*#b)e05%uz</pre> | <pre>{{ randFromStr "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)" 50 }}</pre> |
| `laravel` | <pre>base64:nA80Uh9287Ubq+XM2Ph+lxhhTqc8Y2Jrtv4k86T52no=</pre> | <pre>base64:{{ binary 32 \| b64enc }}</pre> |
| `pin:6` | <pre>458273</pre> | <pre>{{ num . }}</pre> |
| `words:4` | <pre>pushcart unmade stifling avenge</pre> | <pre>{{ words . \| join " " }}</pre> |


### SEE ALSO
* [pwgen](pwgen.md)  - Generate passphrases