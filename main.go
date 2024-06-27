package main

import (
	"github.com/mkideal/cli"
	"lucas.blin/v2/config"
	"os"
)

type ArgT struct {
	cli.Helper
	Recipient  string `cli:"recipient,r" usage:"short and long format flags both are supported"`
	Sender     string `cli:"sender,s" usage:"short and long format flags both are supported"`
	Message    string `cli:"message,m" usage:"short and long format flags both are supported"`
	Secret     string `cli:"secret" usage:"short and long format flags both are supported"`
	Key        string `cli:"key" usage:"short and long format flags both are supported"`
	Contact    string `cli:"add-contact,add" usage:"short and long format flags both are supported"`
	ConfigPath string `cli:"config-path" usage:"short and long format flags both are supported"`
}

func main() {
	os.Exit(cli.Run(new(ArgT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ArgT)
		return RunCli(argv)
	}))
}

func RunCli(argv *ArgT) error {
	if len(argv.ConfigPath) == 0 {
		argv.ConfigPath = "~/.config/parameters.ini"
	}
	config.Init(argv.ConfigPath)

	if len(argv.Sender) > 0 && len(argv.Recipient) > 0 && len(argv.Message) > 0 {
		err := Send(argv.Sender, argv.Recipient, argv.Message, argv.ConfigPath)
		if err != nil {
			return err
		}
		return nil
	}

	if argv.Contact != "" {
		Contact(argv)
		return nil
	}

	if argv.Secret != "" {
		err := configuration("API-SECRET", argv.Secret, argv.ConfigPath)
		if err != nil {
			return err
		}
		config.Green("✅  API-SECRET added to the configuration.")
		return nil
	}

	if argv.Key != "" {
		err := configuration("API-KEY", argv.Key, argv.ConfigPath)
		if err != nil {
			return err
		}
		config.Green("✅  API-KEY added to the configuration.")
		return nil
	}

	return nil
}

func configuration(key string, value string, path string) error {
	err := config.SetConfig("api", key, value, path)
	if err != nil {
		return err
	}
	return nil
}
