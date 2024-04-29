package config

import (
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

func GetConfig(section string, key string, mandatory bool) (string, error) {
	cfg, err := load()
	if err != nil {
		return "", SilentError{}
	}

	Key := cfg.Section(section).Key(key).String()
	if Key == "" && mandatory {
		Red("⚠️ Please set " + key + " in parameters.ini file ⚠️")
		return "", SilentError{}
	}

	return Key, nil
}

func SetConfig(section string, key string, input string) error {
	cfg, err := load()

	config, err := GetConfig(section, key, false)
	if err != nil {
		return err
	}
	if err != nil {
		Red("⚠️ Error getting the configuration: " + err.Error() + "⚠️")
		return err
	} else if config != "" {
		Yellow("⚠️ The key already exists, previous value will be overwritten : " + config + "  ⚠️")
	}

	cfg.Section(section).Key(key).SetValue(input)
	err = cfg.SaveTo(home("~/.config/parameters.ini"))
	if err != nil {
		Red("⚠️ Could not save the configuration file. ⚠️")
		return err
	}
	return nil
}

func load() (*ini.File, error) {
	cfg, err := ini.Load(home("~/.config/parameters.ini"))
	if err != nil {
		Red("⚠️ Could not load the configuration file. ⚠️")
		return nil, SilentError{}
	}
	return cfg, nil
}

func home(path string) string {
	if path[:2] == "~/" {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}
