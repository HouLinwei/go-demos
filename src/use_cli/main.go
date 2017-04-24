package main

import (
	"github.com/urfave/cli"
	"fmt"
	"time"
	"math/rand"
	"os"
	"reflect"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	app := cli.NewApp()
	app.Name = "hello cli"
	app.Usage = "none"
	app.Commands = []cli.Command{
		{
			Name:"test",
			Usage:"None",
			Flags:[]cli.Flag{
				cli.IntSliceFlag{
					Name:"t",
					Value:nil,
					Usage:"None",
				},
			},
			Action:Do,
		},
	}
	app.Run(os.Args)
}


func Do(c *cli.Context)error{
	fmt.Println("run it")
	slice := c.IntSlice("t")
	fmt.Println("type of:", reflect.TypeOf(slice))
	fmt.Println(slice)
	return nil
}