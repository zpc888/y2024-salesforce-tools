package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"lwc-tools-in-go/service"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "lwc",
		Usage: "rename, copy, move and show dependencies, etc",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dir",
				Usage:    "root dir to look up LWCs",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "dependencies",
				Usage: "show dependencies of LWCs",
				Action: func(c *cli.Context) error {
					lwcDir := c.String("dir")
					log.Println("show dependencies of LWCs in", lwcDir)
					return showDependencies(lwcDir)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func showDependencies(dir string) error {
	fmt.Println("show dependencies of LWCs in", dir)
	files, err := service.ListAllFiles(dir, func(path *string, info *os.FileInfo) bool {
		return strings.Contains(*path, "__tests__") ||
			!(strings.HasSuffix(*path, ".js") ||
				strings.HasSuffix(*path, ".html") ||
				strings.HasSuffix(*path, ".css") ||
				strings.HasSuffix(*path, ".js-meta.xml"))
	})
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file)
	}
	lwcComps, err2 := service.ParseLwcComps(files)
	if err2 != nil {
		return err2
	}
	for _, comp := range lwcComps {
		fmt.Println(comp)
	}
	return nil
}
