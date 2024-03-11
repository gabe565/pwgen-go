package config

type Config struct {
	Count    int    `toml:"count" comment:"Number of passphrases to generate."`
	Template string `toml:"template" comment:"Template used to generate passphrases."`
}

func NewDefault() *Config {
	return &Config{
		Count:    10,
		Template: `{{ wordsWithNumber 3 | join "-" | title }}`,
	}
}
