package config

import (
	"github.com/go-ini/ini"
	"io/fs"
	"os"
)

func Init(path string) {
	_, err := ini.Load(home(path))
	if err != nil {
		Green("▶️ Initializing the parameters file at ~/.config/parameters.ini")
		err := os.WriteFile(home(path), []byte{}, fs.ModePerm)
		if err != nil {
			Red("Could not create the parameters file, please check the permissions")
		}
	}
}
