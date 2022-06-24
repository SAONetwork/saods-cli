package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"sao-datastore-cli/ds-cli/config"
)

var cfg config.Config

func main() {
	app := cli.NewApp()
	app.Name = "saods"
	app.Version = "v1.0.0"
	app.Usage = "SAO Data Store"
	app.Action = run
	app.Commands = []*cli.Command{
		{
			Name:  "addFile",
			Usage: "add a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "localPath",
					Usage: "file local path",
				},
				&cli.StringFlag{
					Name:  "appId",
					Usage: "application id which the uploaded object belongs to",
				},
				&cli.StringFlag{
					Name:  "apiKey",
					Usage: "api key used for authentication in REST APIs",
				},
				&cli.BoolFlag{
					Name:  "pretty",
					Usage: "return formatted json response",
				},
			},
			Action: func(c *cli.Context) error {
				return AddFile(c)
			},
		},
		{
			Name:  "getFile",
			Usage: "get a file, you can set either fileId or hash to get your file, if both parameters set, the command will use fileId to get file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "fileId",
					Usage: "id of the uploaded object",
				},
				&cli.StringFlag{
					Name:  "hash",
					Usage: "ipfs hash of the uploaded object",
				},
				&cli.StringFlag{
					Name:  "localPath",
					Usage: "Specify the local path where the file be stored",
				},
				&cli.StringFlag{
					Name:  "appId",
					Usage: "application id which the uploaded object belongs to",
				},
				&cli.StringFlag{
					Name:  "apiKey",
					Usage: "api key used for authentication in REST APIs",
				},
			},
			Action: func(c *cli.Context) error {
				return GetFile(c)
			},
		},
		{
			Name:  "listFiles",
			Usage: "list files, allows user to navigate files by page",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "page",
					Usage: "page number",
				},
				&cli.StringFlag{
					Name:  "size",
					Usage: "page size",
				},
				&cli.StringFlag{
					Name:  "appId",
					Usage: "application id which the uploaded object belongs to",
				},
				&cli.StringFlag{
					Name:  "apiKey",
					Usage: "api key used for authentication in REST APIs",
				},
				&cli.BoolFlag{
					Name:  "pretty",
					Usage: "return formatted json response",
				},
			},
			Action: func(c *cli.Context) error {
				return listFiles(c)
			},
		},
		{
			Name:  "config",
			Usage: "config appId and apiKey, instead of setting the values everytime",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "appId",
					Usage: "application id which the uploaded object belongs to",
				},
				&cli.StringFlag{
					Name:  "apiKey",
					Usage: "api key used for authentication in REST APIs",
				},
				&cli.StringFlag{
					Name:  "serviceUrl",
					Usage: "set service url manually",
				},
			},
			Action: func(c *cli.Context) error {
				return setConfigFile(c)
			},
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	return exec()
}

func exec() error {
	return nil
}
