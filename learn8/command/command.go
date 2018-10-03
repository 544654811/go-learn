package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func testArgs() {
	app := cli.NewApp()
	app.Name = "zzs"
	app.Usage = "my frist app..."

	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
		}
		fmt.Printf("cmd: %s \n", cmd)
		return nil
	}
	app.Run(os.Args)
}

func testFlag() {
	app := cli.NewApp()
	app.Name = "zzs"
	app.Usage = "my frist app..."

	var language string
	var recusive bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "recusive, r",
			Usage:       "recusive for this greeting",
			Destination: &recusive,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("language : %s \n", language)
		fmt.Printf("recusive : %t \n", recusive)
		return nil
	}

	app.Run(os.Args)
}

func Test() {
	// testArgs()
	testFlag()
}
