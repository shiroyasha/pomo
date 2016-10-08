package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func status(c *cli.Context) error {
	fmt.Println("boom! I say!")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "pomo"
	app.Usage = "Pomodoro timer for the Command Line"
	app.Action = status
	app.Run(os.Args)
}
