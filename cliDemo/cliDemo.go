package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func Myapp() {
	//app := &cli.App{
	//	Name: "myapp",
	//	Usage: "make an explosive entrance",
	//	Action: func(c *cli.Context) error {
	//		fmt.Println("boom! I say!", c.Args().Get(0))
	//		return nil
	//	},
	//}

	app := cli.NewApp()
	app.Name = "myapp"
	app.Usage = "cli app demo"
	app.UseShortOptionHandling = true	// -n -l == -nl

	var (
		lang string
		name string
	)

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "lang",
			Value:       "english",                   // default value
			Usage:       "`lang`: language of myapp", // --lang lang, -l lang  lang: language of myapp (default: chinese)
			Destination: &lang,                       // --lang $lang
			Aliases:     []string{"l"},               // --lang or -l
			Required:    false,                       //  Required flag "lang" need set
			DefaultText: "chinese",                   // default value 覆盖 `Value`
		},
		&cli.StringFlag{
			Name:        "name",
			Usage:       "hello `NAME`",
			Value:       "marvin",
			Aliases:     []string{"n"},
			Destination: &name,
			//EnvVars: []string{"GOPATH", "GOROOT"},
		},

	}

	start := func(c *cli.Context) error {
		fmt.Println("---- cli app demo ----")
		//fmt.Println("c.Args: ", c.Args().Get(0))
		//fmt.Println("os.Args: ", os.Args)

		//if c.NArg() > 0 {
		//	name = c.Args().Get(0)
		//}
		//fmt.Println(c.String("lang"))

		//if c.String("lang") == "english" {
		//	fmt.Println("hello", name)
		//} else {
		//	fmt.Println("nihao", name)
		//}

		fmt.Println("language:", lang)
		fmt.Println("name:", name)
		if lang == "english" {
			fmt.Println("hello", name)
		} else {
			fmt.Println("nihao", name)
		}

		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Action:  start,
			Subcommands: []*cli.Command{
				{
					Name: "sub",
					Action: func(c *cli.Context) error {
						fmt.Println("start sub command", name)
						return nil
					},
					Flags: []cli.Flag{		// subCommand flag
						&cli.StringFlag{
							Name:        "name",
							Usage:       "hello `NAME`",
							Value:       "marvin",
							Aliases:     []string{"n"},
							Destination: &name,
							//EnvVars: []string{"GOPATH", "GOROOT"},
						},
					},
				},
			},
		},
	}

	app.Action = start

	sort.Sort(cli.FlagsByName(app.Flags)) // ASC sort
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Myapp()
}
