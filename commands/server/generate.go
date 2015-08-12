package server

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/certm/utils"
	"github.com/ehazlett/tlsutils"
)

func generate(c *cli.Context) {
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
	if err := utils.CreateIfNotExists(certPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// key
	if err := utils.CreateIfNotExists(keyPath, overwrite); err != nil {
		log.Fatal(err)
	}

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
