package config

import (
	"github.com/go-ini/ini"
	"io/fs"
	"os"
)

func Init() {
	_, err := ini.Load(home("~/.config/parameters.ini"))
	if err != nil {
		Green("▶️ Initializing the parameters file at ~/.config/parameters.ini")
		err := os.WriteFile(home("~/.config/parameters.ini"), []byte{}, fs.ModePerm)
		if err != nil {
			Red("Could not create the parameters file, please check the permissions")
		}
	}
}
