package main

import (
	"net/url"

	"github.com/deviceplane/deviceplane/pkg/client"
	"github.com/urfave/cli"
)

var (
	projectFlag = cli.StringFlag{
		Name:   "project",
		EnvVar: "DEVICEPLANE_PROJECT",
	}
	applicationFlag = cli.StringFlag{
		Name:   "application",
		EnvVar: "DEVICEPLANE_APPLICATION",
	}
	deviceFlag = cli.StringFlag{
		Name:   "device",
		EnvVar: "DEVICEPLANE_DEVICE",
	}
)

func withClient(c *cli.Context, f func(*client.Client) error) error {
	u, err := url.Parse(c.GlobalString("url"))
	if err != nil {
		return err
	}
	accessKey := c.GlobalString("access-key")
	return f(client.NewClient(u, accessKey, nil))
}
