# pwgen-go
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/gabe565/pwgen-go)](https://github.com/gabe565/pwgen-go/releases)
[![Build](https://github.com/gabe565/pwgen-go/actions/workflows/build.yaml/badge.svg)](https://github.com/gabe565/pwgen-go/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/pwgen-go)](https://goreportcard.com/report/github.com/gabe565/pwgen-go)

Command line passphrase generator written in Go.

The [EFF Diceware Wordlists](https://www.eff.org/dice) are embedded, with the long wordlist used by default. See the [EFF's Deep Dive](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) for more details on the benefits of this word list.

## Installation

### APT (Ubuntu, Debian)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo apt install ca-certificates
   ```

2. Add gabe565 apt repository
   ```
   echo 'deb [trusted=yes] https://apt.gabe565.com /' | sudo tee /etc/apt/sources.list.d/gabe565.list
   ```

3. Update apt repositories
   ```shell
   sudo apt update
   ```

4. Install pwgen-go
   ```shell
   sudo apt install pwgen-go
   ```
</details>

### RPM (CentOS, RHEL)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo dnf install ca-certificates
   ```

2. Add gabe565 rpm repository to `/etc/yum.repos.d/gabe565.repo`
   ```ini
   [gabe565]
   name=gabe565
   baseurl=https://rpm.gabe565.com
   enabled=1
   gpgcheck=0
   ```

3. Install pwgen-go
   ```shell
   sudo dnf install pwgen-go
   ```
</details>

### AUR (Arch Linux)

<details>
  <summary>Click to expand</summary>

Install [pwgen-go-bin](https://aur.archlinux.org/packages/pwgen-go-bin) with your [AUR helper](https://wiki.archlinux.org/index.php/AUR_helpers) of choice.
</details>

### Homebrew (macOS, Linux)

<details>
  <summary>Click to expand</summary>

Install pwgen-go from [gabe565/homebrew-tap](https://github.com/gabe565/homebrew-tap):
```shell
brew install gabe565/tap/pwgen-go
```
</details>

### Manual Installation

<details>
  <summary>Click to expand</summary>

Download and run the [latest release binary](https://github.com/gabe565/pwgen-go/releases/latest) for your system and architecture.
</details>

## Usage

Run `pwgen` to generate a list of random passphrases.

Predefined profiles are available — see [`pwgen profiles`](docs/pwgen_profiles.md) for a list of defaults. You can specify a profile or a template as a positional argument. If the argument resembles a Go template, it will be interpreted as such; otherwise, it will be treated as a profile. (For explicit control, you can also use the `--template` (`-t`) or `--profile` (`-p`) flags.)

To change the number of generated results, pass a number to the `--count` (`-c`) flag.

Pwgen takes advantage of shell completion; for example, when choosing a profile, try `pwgen <TAB><TAB>`.

Also see the generated [docs](docs/pwgen.md).

### Examples

By default, pwgen will generate passwords with 4 words from the [EFF Long Wordlist](https://www.eff.org/dice) and a random number:
```shell
$ pwgen
Pebbly1-Destitute-Shelve-Renter
Mushiness-Possibly0-Trustee-Outfit
Trimness-Freebase-Spotless-Manmade2
Recall7-Uncrushed-Agreeing-Crepe
Tavern-Moustache1-Scrubber-Mashed
Lyrically2-Comic-Imitate-Victory
Idiom-Jockey-Subheader6-Scallop
Blot-Sympathy-Nurture-Spotter5
Womanhood6-Capsize-Zigzagged-Siesta
Unwomanly-Unwashed-Urchin-Empty1
```
#### Profiles

```shell
$ pwgen -c1 alphanum:16
cFAutu4OizPQ0d3N
$ pwgen -c1 words:5
tinker coagulant bundle gave deviator
$ pwgen -c1 pin:6
771283
```
#### Templates

```shell
$ pwgen -c1 '{{ binary 16 | b64enc }}'
kENvfNLHw/GcwID3TGEIgg==
```

## Configuration

A configuration file will be generated the first time pwgen-go is run. Depending on your operating system, the file will be available at:
- **Windows:** `%AppData%\pwgen-go\config.toml`
- **macOS:** `~/Library/Application Support/pwgen-go/config.toml`
- **Linux:** `~/.config/pwgen-go/config.toml`

An example configuration is also available at [`config_example.toml`](config_example.toml).

## Templating

Templated passphrases are generated using Go's [text/template](https://pkg.go.dev/text/template) package.

All [Sprig functions](https://masterminds.github.io/sprig/) are available, plus some extras documented [here](docs/pwgen_functions.md).

## Security

- **Local Generation:** All passphrases are generated directly on your device. No data is sent or received externally.
- **Bundled Resources:** The EFF wordlists are embedded within the binary, ensuring no external data is fetched during generation.
- **Cryptographic Quality:** All random functions use cryptographically secure (`crypto/rand`) random strings.
- **Anonymity by Design:** By default, 10 passphrases are generated to help obscure which one you choose.
