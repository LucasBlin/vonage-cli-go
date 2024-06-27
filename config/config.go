package config

import (
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

var Url = "https://rest.nexmo.com/sms/json"

func GetConfig(section string, key string, mandatory bool, path string) (string, error) {
	cfg, err := load(path)
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

func SetConfig(section string, key string, input string, path string) error {
	cfg, err := load(path)

	config, err := GetConfig(section, key, false, path)
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
	err = cfg.SaveTo(home(path))
	if err != nil {
		Red("⚠️ Could not save the configuration file. ⚠️")
		return err
	}
	return nil
}

func load(path string) (*ini.File, error) {
	cfg, err := ini.Load(home(path))
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
