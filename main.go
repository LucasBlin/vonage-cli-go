package main

import (
	"github.com/mkideal/cli"
	"lucas.blin/v2/cmd"
	"lucas.blin/v2/cmd/app/send"
	"lucas.blin/v2/config"
	"os"
	"strings"
)

type argT struct {
	cli.Helper
	Recipient string `cli:"recipient,r" usage:"short and long format flags both are supported"`
	Sender    string `cli:"sender,s" usage:"short and long format flags both are supported"`
	Message   string `cli:"message,m" usage:"short and long format flags both are supported"`
	Secret    string `cli:"secret" usage:"short and long format flags both are supported"`
	Key       string `cli:"key" usage:"short and long format flags both are supported"`
	Contact   string `cli:"add-contact,add" usage:"short and long format flags both are supported"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		config.Init()
		if len(argv.Sender) > 0 && len(argv.Recipient) > 0 && len(argv.Message) > 0 {
			err := send.Send(argv.Sender, argv.Recipient, argv.Message)
			if err != nil {
				return err
			}
		}

		if argv.Contact != "" {
			contact(argv)
			return nil
		}

		if argv.Secret != "" {
			err2 := configuration("API-SECRET", argv.Secret)
			if err2 != nil {
				return err2
			}
			config.Green("✅  API-SECRET added to the configuration.")
		}

		if argv.Key != "" {
			err3 := configuration("API-KEY", argv.Key)
			if err3 != nil {
				return err3
			}
			config.Green("✅  API-KEY added to the configuration.")
		}

		return nil
	}))
}

func configuration(key string, value string) error {
	err := config.SetConfig("api", key, value)
	if err != nil {
		return err
	}
	return nil
}

func contact(argv *argT) {
	c := strings.Split(argv.Contact, ",")
	if cmd.ValidNumber(c[1]) == true {
		if err := config.SetConfig("contacts", c[0], c[1]); err != nil {
			config.Red("❌  Fail to add the contact to the configuration.")
		}
		config.Green("✅  Contact added to the configuration.")
	} else {
		config.Red("❌  Invalid number, please provide it as : 336XXXXXXXX")
	}
}
