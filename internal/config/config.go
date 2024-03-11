package config

type Config struct {
	Count    int    `toml:"count" comment:"Number of passphrases to generate."`
	Template string `toml:"template" comment:"Template used to generate passphrases."`
	Wordlist string `toml:"wordlist" comment:"Wordlist to use. (one of: long, short1, short2)"`
}

const (
	WordlistLong   = "long"
	WordlistShort1 = "short1"
	WordlistShort2 = "short2"
)

func NewDefault() *Config {
	return &Config{
		Count:    10,
		Template: `{{ wordsWithNumber 3 | join "-" | title }}`,
		Wordlist: WordlistLong,
	}
}
