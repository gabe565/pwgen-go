package main

import (
	"os"

	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	f, err := os.Create("config_example.toml")
	if err != nil {
		panic(err)
	}

	encoder := toml.NewEncoder(f)
	conf := config.NewDefault()
	if err := encoder.Encode(conf); err != nil {
		panic(err)
	}
}
