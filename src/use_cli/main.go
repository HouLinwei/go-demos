package main

import (
	"fmt"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app := cli.NewApp()
	app.Name = "hello cli"
	app.Usage = "none"
	app.Commands = []cli.Command{
		{
			Name:  "test",
			Usage: "None",
			Flags: []cli.Flag{
				cli.IntSliceFlag{
					Name:  "t",
					Value: nil,
					Usage: "None",
				},
			},
			Action: Do,
		},
	}
	app.Run(os.Args)
}

func Do(c *cli.Context) error {
	fmt.Println("run it")
	slice := c.IntSlice("t")
	fmt.Println("type of:", reflect.TypeOf(slice))
	fmt.Println(slice)
	return nil
}
