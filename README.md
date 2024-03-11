# pwgen-go

Command line passphrase generator written in Go.

The [EFF Long Wordlist](https://www.eff.org/dice) is embedded which includes 7776 words. See the [EFF's Deep Dive](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) for more details on the benefits of this word list.

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

## Templating

Templated passphrases are generated using Go's [text/template](https://pkg.go.dev/text/template) package.

All [Sprig functions](https://masterminds.github.io/sprig/) are available, plus some extras listed below.

### Functions

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

#### `wordsWithNumber`

Behaves similarly to [`words`](#words), but will append a random number to one of the words.

##### Examples
```gotemplate
{{ wordsWithNumber 3 | join "-" }}
```

#### `number`, `num`

Alias for [Sprig's `randNumeric` function](https://masterminds.github.io/sprig/strings.html#randalphanum-randalpha-randnumeric-and-randascii). A random number will be generated with the number of digits determined by the parameter.

##### Examples
- Generate a number from 0-9:
  ```gotemplate
  {{ number 1 }}
  ```
- Generate a number from 10-99:
  ```gotemplate
  {{ number 2 }}
  ```
