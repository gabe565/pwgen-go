//go:build !wasm

package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

func flagToConfigMapping() map[string]string {
	return map[string]string{}
}

var ErrProfileNotFound = errors.New("profile not found")

func Load(cmd *cobra.Command, save bool) (*Config, error) {
	k := koanf.New(".")
	conf := NewDefault()

	// Load default config
	if err := k.Load(structs.Provider(conf, "toml"), nil); err != nil {
		return nil, err
	}

	// Find config file
	cfgFile, err := cmd.Flags().GetString(FlagConfig)
	if err != nil {
		return nil, err
	}
	if cfgFile == "" {
		if cfgFile, err = GetFile(); err != nil {
			return nil, err
		}
	}

	var cfgNotExists bool
	// Load config file if exists
	cfgContents, err := os.ReadFile(cfgFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfgNotExists = true
		} else {
			return nil, err
		}
	}

	// Parse config file
	parser := TOMLParser{}
	if err := k.Load(rawbytes.Provider(cfgContents), parser); err != nil {
		return nil, err
	}

	if save {
		if err := k.UnmarshalWithConf("", conf, koanf.UnmarshalConf{Tag: "toml"}); err != nil {
			return nil, err
		}

		newCfg, err := toml.Marshal(conf)
		if err != nil {
			return nil, err
		}

		if !bytes.Equal(cfgContents, newCfg) {
			if cfgNotExists {
				if err := os.MkdirAll(filepath.Dir(cfgFile), 0o777); err != nil {
					return nil, err
				}
			}

			if err := os.WriteFile(cfgFile, newCfg, 0o666); err != nil {
				return nil, err
			}
		}
	}

	// Load flags
	flagMapping := flagToConfigMapping()
	err = k.Load(posflag.ProviderWithValue(cmd.Flags(), ".", k, func(key string, value string) (string, any) {
		if k, ok := flagMapping[key]; ok {
			key = k
		}
		return key, value
	}), nil)
	if err != nil {
		return nil, err
	}

	if err := k.UnmarshalWithConf("", conf, koanf.UnmarshalConf{Tag: "toml"}); err != nil {
		return nil, err
	}

	if conf.Profile.Name != "" && (conf.Template == "" || cmd.Flags().Lookup(FlagProfile).Changed && !cmd.Flags().Lookup(FlagTemplate).Changed) {
		profile, ok := conf.Profiles[conf.Profile.Name]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrProfileNotFound, conf.Profile.Name)
		}
		conf.Template = profile.Template
		if conf.Profile.Param != 0 {
			conf.Param = conf.Profile.Param
		} else {
			conf.Param = profile.Param
		}
	}

	if !strings.HasSuffix(conf.Template, "\n") {
		conf.Template += "\n"
	}

	return conf, err
}
