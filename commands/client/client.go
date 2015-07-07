package client

import (
	"github.com/codegangsta/cli"
)

var CmdClient = cli.Command{
	Name:  "client",
	Usage: "client certificate management",
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
			Name:  "ca-cert",
			Value: "",
			Usage: "CA certificate for signing (defaults to ca.pem in output dir)",
		},
		cli.StringFlag{
			Name:  "ca-key",
			Value: "",
			Usage: "CA key for signing (defaults to ca-key.pem in output dir)",
		},
		cli.StringFlag{
			Name:  "cert",
			Value: "",
			Usage: "certificate name (default: server.pem)",
		},
		cli.StringFlag{
			Name:  "key",
			Value: "",
			Usage: "key name (default: server-key.pem)",
		},
		cli.StringFlag{
			Name:  "common-name, c",
			Value: "",
			Usage: "common name",
		},
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
