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
Run `pwgen` for a list of generated passphrases. The template can be customized with the `--template` (`-t`) flag, and the number of generated entries can be customized with `--count` (`-c`).  
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

All [Sprig functions](https://masterminds.github.io/sprig/) are available, plus some extras listed below.

### Functions

- [`word`](#word)
- [`words`](#words)
- [`wordsWithNum`](#wordswithnum-wordswithnumber)
- [`num`](#num-number-numeric)
- [`alpha`](#alpha)
- [`alphaNum`](#alphaNum)
- [`ascii`](#ascii)
- [`binary`](#binary)

#### `word`

Outputs a random word from the wordlist. For title case, the output can be piped to `title`.

##### Examples
- Lowercase:
  ```gotemplate
  {{ word }}
  ```
- Title case:
  ```gotemplate
  {{ title word }}
  ```

#### `words`

Outputs a slice of random words from the wordlist. The output will be a slice, which can be joined using `join`. For title case, the output can be piped to `title`.

##### Examples
- Lowercase, joined with `-`:
  ```gotemplate
  {{ words 3 | join "-" }}
  ```
- Title case, joined with `-`:
  ```gotemplate
  {{ words 3 | join "-" | title }}
  ```

#### `wordsWithNum`, `wordsWithNumber`

Behaves similarly to [`words`](#words), but will append a random number to one of the words.

##### Examples
```gotemplate
{{ wordsWithNumber 3 | join "-" }}
```

#### `num`, `number`, `numeric`

Alias for [Sprig's `randNumeric` function](https://masterminds.github.io/sprig/strings.html#randalphanum-randalpha-randnumeric-and-randascii). A random number will be generated with the number of digits determined by the parameter.

##### Examples
- Generate a number from 0-9:
  ```gotemplate
  {{ num 1 }}
  ```
- Generate a number from 10-99:
  ```gotemplate
  {{ num 2 }}
  ```

#### `alpha`

Alias for [Sprig's `randAlpha` function](https://masterminds.github.io/sprig/strings.html#randalphanum-randalpha-randnumeric-and-randascii). Random letters will be generated with the length determined by the parameter.

##### Examples
- Generate a random string of letters:
  ```gotemplate
  {{ alpha 32 }}
  ```

#### `alphaNum`

Alias for [Sprig's `randAlphaNum` function](https://masterminds.github.io/sprig/strings.html#randalphanum-randalpha-randnumeric-and-randascii). Random letters and numbers will be generated with the length determined by the parameter.

##### Examples
- Generate a random string of letters and numbers:
  ```gotemplate
  {{ alphaNum 32 }}
  ```

#### `ascii`

Alias for [Sprig's `randAscii` function](https://masterminds.github.io/sprig/strings.html#randalphanum-randalpha-randnumeric-and-randascii). Random letters, numbers, and symbols will be generated with the length determined by the parameter.

##### Examples
- Generate a random string of letters, numbers, and symbols:
  ```gotemplate
  {{ ascii 32 }}
  ```

#### `binary`

Random binary data will be generated with number of bytes determined by the parameter. Useful with [`b64enc`](https://masterminds.github.io/sprig/encoding.html).

##### Examples
- Generate a random base64 string with 32 bytes of data:
  ```gotemplate
  {{ binary 32 | b64enc }}
  ```
