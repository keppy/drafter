package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "drafter"
	app.Usage = "draft releases with ease"

	app.Action = func(c *cli.Context) error {
		user := c.Args().Get(0)
		tag := c.Args().Get(1)
		description := c.Args().Get(2)

		if len(user) == 0 {
			log.Fatal("You must specify a user")
		}
		if len(tag) == 0 {
			log.Fatal("You must specify a tag")
		}
		if len(description) == 0 {
			log.Fatal("You must specify a release description")
		}

		fmt.Printf("Drafting release %v as %v", tag, user)

		tagCommand := fmt.Sprintf("git tag -a %v -m %v", tag, description)
		fmt.Printf("running: %v", tagCommand)

		err := app.Run(strings.Fields(tagCommand))
		if err != nil {
			log.Fatal(err)
		}

		pushTag := fmt.Sprintf("git push origin %v", tag)
		fmt.Printf("running: %v", pushTag)

		err = app.Run(strings.Fields(pushTag))
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
