package bundle

import (
	"github.com/codegangsta/cli"
)

var CmdBundle = cli.Command{
	Name:  "bundle",
	Usage: "generate CA, server and client certs",
	Subcommands: []cli.Command{
		cmdGenerate,
	},
}

var cmdGenerate = cli.Command{
	Name:   "generate",
	Usage:  "generate new bundle",
	Action: generate,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "host",
			Value: &cli.StringSlice{},
			Usage: "SAN/IP SAN for certificate",
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
