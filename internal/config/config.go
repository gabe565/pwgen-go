package config

type Config struct {
	Count     int               `toml:"count" comment:"Number of passphrases to generate."`
	Template  string            `toml:"template" comment:"Default template used to generate passphrases."`
	Templates map[string]string `toml:"templates" comment:"Named templates"`
	Wordlist  string            `toml:"wordlist" comment:"Wordlist to use. (one of: long, short1, short2)"`
}

const (
	WordlistLong   = "long"
	WordlistShort1 = "short1"
	WordlistShort2 = "short2"
)

func NewDefault() *Config {
	return &Config{
		Count:    10,
		Template: "diceware-3",
		Templates: map[string]string{
			"alpha-16":   "{{ randAlpha 16 }}",
			"alpha-32":   "{{ randAlpha 32 }}",
			"alpha-64":   "{{ randAlpha 64 }}",
			"ascii-16":   "{{ randAscii 16 }}",
			"ascii-32":   "{{ randAscii 32 }}",
			"ascii-64":   "{{ randAscii 64 }}",
			"diceware-3": `{{ wordsWithNumber 3 | join "-" | title }}`,
			"diceware-5": `{{ wordsWithNumber 5 | join "-" | title }}`,
			"diceware-6": `{{ wordsWithNumber 6 | join "-" | title }}`,
			"pin-4":      "{{ num 4 }}",
			"pin-6":      "{{ num 6 }}",
		},
		Wordlist: WordlistLong,
	}
}
