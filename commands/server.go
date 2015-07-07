package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/tlsutils"
)

var CmdServer = cli.Command{
	Name:  "server",
	Usage: "server certificate management",
	Subcommands: []cli.Command{
		cmdServerGenerate,
	},
}

var cmdServerGenerate = cli.Command{
	Name:   "generate",
	Usage:  "generate new certificate",
	Action: serverGenerate,
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

func serverGenerate(c *cli.Context) {
	outputDir := c.GlobalString("output-directory")
	if outputDir == "" {
		cli.ShowAppHelp(c)
		log.Fatalf("you must specify an output directory")
	}

	hosts := c.StringSlice("host")
	caCertPath := c.String("ca-cert")
	caKeyPath := c.String("ca-key")
	certPath := c.String("cert")
	keyPath := c.String("key")
	org := c.String("org")
	bits := c.Int("bits")
	overwrite := c.Bool("overwrite")

	if caCertPath == "" {
		caCertPath = filepath.Join(outputDir, "ca.pem")
	}

	if caKeyPath == "" {
		caKeyPath = filepath.Join(outputDir, "ca-key.pem")
	}

	if certPath == "" {
		certPath = filepath.Join(outputDir, "server.pem")
	}

	if keyPath == "" {
		keyPath = filepath.Join(outputDir, "server-key.pem")
	}

	log.Printf("generating server certificate: org=%s bits=%d",
		org,
		bits,
	)

	log.Debugf("creating output dir: path=%s", outputDir)
	if err := os.MkdirAll(outputDir, 0700); err != nil {
		log.Fatal(err)
	}

	// cert
	fc, err := os.Stat(certPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	if fc != nil {
		if !overwrite {
			log.Fatalf("cert/key exists.  specify --overwrite to overwrite.")
		}

		if err := os.Remove(certPath); err != nil {
			log.Fatal(err)
		}
	}

	crt, err := os.Create(certPath)
	if err != nil {
		log.Fatal(err)
	}
	crt.Close()

	// key
	fk, err := os.Stat(keyPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	if fk != nil {
		if !overwrite {
			log.Fatalf("cert/key exists.  specify --overwrite to overwrite.")
		}

		if err := os.Remove(keyPath); err != nil {
			log.Fatal(err)
		}
	}

	k, err := os.Create(keyPath)
	if err != nil {
		log.Fatal(err)
	}
	k.Close()

	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatal(err)
	}

	caKey, err := ioutil.ReadFile(caKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	cert, key, err := tlsutils.GenerateCertificate(
		hosts,
		caCert,
		caKey,
		org,
		"",
		bits,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating certificate: path=%s", certPath)
	if err := ioutil.WriteFile(certPath, cert, 0600); err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating key: path=%s", keyPath)
	if err := ioutil.WriteFile(keyPath, key, 0600); err != nil {
		log.Fatal(err)
	}
}
