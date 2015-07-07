package ca

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
	if err := utils.CreateIfNotExists(caCertPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// key
	if err := utils.CreateIfNotExists(caKeyPath, overwrite); err != nil {
		log.Fatal(err)
	}

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
