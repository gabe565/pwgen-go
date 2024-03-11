package config

type Config struct {
	Count    int    `toml:"count" comment:"Number of passwords to generate."`
	Template string `toml:"template" comment:"Template used to generate passwords."`
}

func NewDefault() *Config {
	return &Config{
		Count:    10,
		Template: `{{ words 3 | join "-" | title }}{{ number 1 }}`,
	}
}
