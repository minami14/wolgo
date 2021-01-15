package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/minami14/wolgo/wol"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wolgo"
	app.Usage = "Wake on LAN"
	app.UsageText = fmt.Sprintf("%v [target mac address]", os.Args[0])
	app.Action = func(c *cli.Context) error {
		if !c.Args().Present() {
			return errors.New("wolgo requires 1 argument. it is mac address")
		}
		err := wol.WakeOnLan(c.Args().Get(0))
		return err
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
