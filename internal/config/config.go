package config

import (
	"strconv"
	"strings"

	"gabe565.com/pwgen/internal/wordlist"
)

type Config struct {
	File string `toml:"-"`

	Count    int           `toml:"count"    comment:"Number of passphrases to generate."`
	Profile  ProfileRef    `toml:"profile"  comment:"Default profile used to generate passphrases."`
	Param    any           `toml:"-"`
	Profiles ProfileMap    `toml:"profiles" comment:"Preconfigured profiles and default parameters."`
	Wordlist wordlist.Meta `toml:"wordlist" comment:"Wordlist to use. (one of: long, short1, short2)"`
	Template string        `toml:"template" comment:"Default template used to generate passphrases. If not empty, will override the default profile."`
}

type ProfileMap map[string]Profile

func (p ProfileMap) Named() []NamedProfile {
	result := make([]NamedProfile, 0, len(p))
	for k, v := range p {
		result = append(result, NamedProfile{k, v})
	}
	return result
}

type Profile struct {
	Template string `toml:"template"`
	Param    int    `toml:"param,omitempty"`
}

type NamedProfile struct {
	Name string
	Profile
}

type ProfileRef struct {
	Name  string
	Param int
}

func (p *ProfileRef) MarshalText() ([]byte, error) {
	if p.Param == 0 {
		return []byte(p.Name), nil
	}
	return []byte(p.Name + ":" + strconv.Itoa(p.Param)), nil
}

func (p *ProfileRef) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ":")
	if len(parts) >= 1 {
		p.Name = parts[0]
		if len(parts) >= 2 && parts[1] != "" {
			parsed, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			p.Param = parsed
		}
	}
	return nil
}

func New() *Config {
	return &Config{
		Count:   10,
		Profile: ProfileRef{"diceware", 4},
		Profiles: map[string]Profile{
			"alpha":    {"{{ alpha . }}", 32},
			"alphanum": {"{{ alphaNum . }}", 32},
			"base64":   {"{{ binary . | b64enc }}", 32},
			"ascii":    {"{{ ascii . }}", 32},
			"diceware": {`{{ wordsWithNum . | join "-" | title }}`, 4},
			"django":   {`{{ randFromStr "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)" 50 }}`, 0},
			"laravel":  {`base64:{{ binary 32 | b64enc }}`, 0},
			"pin":      {"{{ num . }}", 6},
			"words":    {`{{ words . | join " " }}`, 4},
		},
		Wordlist: wordlist.Long,
	}
}
