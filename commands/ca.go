package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/tlsutils"
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
	Action: caGenerate,
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

func caGenerate(c *cli.Context) {
	outputDir := c.GlobalString("output-directory")
	if outputDir == "" {
		cli.ShowAppHelp(c)
		log.Fatalf("you must specify an output directory")
	}

	caCertPath := filepath.Join(outputDir, "ca.pem")
	caKeyPath := filepath.Join(outputDir, "ca-key.pem")
	org := c.String("org")
	bits := c.Int("bits")
	overwrite := c.Bool("overwrite")

	log.Printf("generating ca: org=%s bits=%d", org, bits)

	log.Debugf("creating output dir: path=%s", outputDir)
	if err := os.MkdirAll(outputDir, 0700); err != nil {
		log.Fatal(err)
	}

	// cert
	fc, err := os.Stat(caCertPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	if fc != nil {
		if !overwrite {
			log.Fatalf("ca cert/key exists.  specify --overwrite to overwrite.")
		}

		if err := os.Remove(caCertPath); err != nil {
			log.Fatal(err)
		}
	}

	crt, err := os.Create(caCertPath)
	if err != nil {
		log.Fatal(err)
	}
	crt.Close()

	// key
	fk, err := os.Stat(caKeyPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	if fk != nil {
		if !overwrite {
			log.Fatalf("ca cert/key exists.  specify --overwrite to overwrite.")
		}

		if err := os.Remove(caKeyPath); err != nil {
			log.Fatal(err)
		}
	}

	k, err := os.Create(caKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	k.Close()

	cert, key, err := tlsutils.GenerateCACertificate(org, bits)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating certificate: path=%s", caCertPath)
	if err := ioutil.WriteFile(caCertPath, cert, 0600); err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating key: path=%s", caKeyPath)
	if err := ioutil.WriteFile(caKeyPath, key, 0600); err != nil {
		log.Fatal(err)
	}
}
