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
Run `pwgen` to generate a list of random passphrases. All random functions use cryptographically secure (`crypto/rand`) random strings.

Predefined profiles can be used with the `--profile` (`-p`) flag. See [`pwgen profiles`](docs/pwgen_profiles.md) for a list of defaults.

Alternatively, the template can be directly customized with the `--template` (`-t`) flag.

To change the number of generated results, pass a different number to the `--count` (`-c`) flag.

Also see the generated [docs](docs/pwgen.md).

### Example
```shell
$ pwgen
Pebbly1-Destitute-Shelve
Mushiness-Possibly0-Trustee
Trimness-Freebase2-Spotless
Recall7-Uncrushed-Agreeing
Tavern-Moustache1-Scrubber
Lyrically2-Comic-Imitate
Idiom-Jockey-Subheader6
Blot-Sympathy9-Nurture
Womanhood6-Capsize-Zigzagged
Unwomanly-Unwashed1-Urchin
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
