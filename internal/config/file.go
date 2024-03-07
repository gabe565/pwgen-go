package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var configDir = "pwgen-go"

func GetDir() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
			return filepath.Join(xdgConfigHome, configDir), nil
		}
		fallthrough
	default:
		dir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}

		dir = filepath.Join(dir, configDir)
		return dir, nil
	}
}

func GetFile() (string, error) {
	dir, err := GetDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.toml"), nil
}

func GetFilePretty() (string, error) {
	file, err := GetFile()
	if err != nil {
		return "", err
	}

	home, _ := os.UserHomeDir()
	return strings.Replace(file, home, "$HOME", 1), nil
}
