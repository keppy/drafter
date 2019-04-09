package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "drafter"
	app.Usage = "draft releases with ease"

	// TODO: use the full template override
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Println("USAGE:")
		fmt.Println("	drafter [user] [git-tag] [description] [deployScriptPath] [IP]")
	}

	app.Action = func(c *cli.Context) error {
		user := c.Args().Get(0)
		tag := c.Args().Get(1)
		description := c.Args().Get(2)
		deployScriptPath := c.Args().Get(3)
		ip := c.Args().Get(4)

		if len(user) == 0 {
			log.Fatal("You must specify a user")
		}
		if len(tag) == 0 {
			log.Fatal("You must specify a tag")
		}
		if len(description) == 0 {
			log.Fatal("You must specify a description")
		}
		if len(deployScriptPath) == 0 {
			log.Fatal("You must specify a deploy script path")
		}
		if len(ip) == 0 {
			log.Fatal("You must specify a deploy target IP")
		}

		fmt.Printf("Drafting release %v as %v", tag, user)

		tagCommand := fmt.Sprintf("git tag -a %v -m %v", tag, description)
		fmt.Printf("running: %v", tagCommand)

		cmd := exec.Command("git", "tag", "-a", tag, "-m", description)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		pushTag := fmt.Sprintf("git push origin %v", tag)
		fmt.Printf("running: %v", pushTag)

		cmd = exec.Command("git", "push", "origin", tag)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		cmd = exec.Command(deployScriptPath, ip, tag)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		// TODO: Use github API to create a release
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
