# pwgen-go

Command line password generator written in Go.

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
Run `pwgen template` for a list of generated passwords. The template can be customized with the `--template` (`-t`) flag, and the number of generated entries can be customized with `--count` (`-c`).  
Also see the generated [docs](docs/pwgen.md).

### Example
```shell
$ pwgen template
Ranch-Advert-Sufficiently6
Flavor-Development-Livecam5
Correctly-Arkansas-Shaw8
Plants-Conducting-Bradford8
Nickname-Costs-Translate6
Meanwhile-Letting-Arlington3
Elimination-Chronic-Chip3
Fundamental-Marcus-Mathematical6
Reports-Sentence-Purchase2
Unlimited-Holland-Shaved3
```
