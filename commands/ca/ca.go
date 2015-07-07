package ca

import (
	"github.com/codegangsta/cli"
)

var CmdCA = cli.Command{
	Name:  "ca",
	Usage: "CA certificate management",
	Subcommands: []cli.Command{
		cmdGenerate,
	},
}

var cmdGenerate = cli.Command{
	Name:   "generate",
	Usage:  "generate new certificate",
	Action: generate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "org, o",
			Value: "unknown",
			Usage: "organization",
		},
		cli.IntFlag{
			Name:  "bits, b",
			Value: 2048,
			Usage: "number of bits in the key (default: 2048)",
		},
		cli.BoolFlag{
			Name:  "overwrite",
			Usage: "overwrite existing certificates and keys",
		},
	},
}
