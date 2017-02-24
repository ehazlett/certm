package bundle

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
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
	caCertPath := filepath.Join(outputDir, "ca.pem")
	caKeyPath := filepath.Join(outputDir, "ca-key.pem")
	serverCertPath := filepath.Join(outputDir, "server.pem")
	serverKeyPath := filepath.Join(outputDir, "server-key.pem")
	clientCertPath := filepath.Join(outputDir, "cert.pem")
	clientKeyPath := filepath.Join(outputDir, "key.pem")
	org := c.String("org")
	bits := c.Int("bits")
	overwrite := c.Bool("overwrite")

	log.Printf("generating ca: org=%s bits=%d", org, bits)

	log.Debugf("creating output dir: path=%s", outputDir)
	if err := os.MkdirAll(outputDir, 0700); err != nil {
		log.Fatal(err)
	}

	// ca cert
	if err := utils.CreateIfNotExists(caCertPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// ca key
	if err := utils.CreateIfNotExists(caKeyPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// server cert
	if err := utils.CreateIfNotExists(serverCertPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// server key
	if err := utils.CreateIfNotExists(serverKeyPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// client cert
	if err := utils.CreateIfNotExists(clientCertPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// client key
	if err := utils.CreateIfNotExists(clientKeyPath, overwrite); err != nil {
		log.Fatal(err)
	}

	// generate ca
	caCert, caKey, err := tlsutils.GenerateCACertificate(org, bits)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating ca certificate: path=%s", caCertPath)
	if err := ioutil.WriteFile(caCertPath, caCert, 0600); err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating ca key: path=%s", caKeyPath)
	if err := ioutil.WriteFile(caKeyPath, caKey, 0600); err != nil {
		log.Fatal(err)
	}

	// generate client
	clientCert, clientKey, err := tlsutils.GenerateCertificate(
		nil,
		caCert,
		caKey,
		org,
		"",
		bits,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating client certificate: path=%s", clientCertPath)
	if err := ioutil.WriteFile(clientCertPath, clientCert, 0600); err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating client key: path=%s", clientKeyPath)
	if err := ioutil.WriteFile(clientKeyPath, clientKey, 0600); err != nil {
		log.Fatal(err)
	}

	// generate server
	serverCert, serverKey, err := tlsutils.GenerateCertificate(
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

	log.Debugf("creating server certificate: path=%s", serverCertPath)
	if err := ioutil.WriteFile(serverCertPath, serverCert, 0600); err != nil {
		log.Fatal(err)
	}

	log.Debugf("creating server key: path=%s", serverKeyPath)
	if err := ioutil.WriteFile(serverKeyPath, serverKey, 0600); err != nil {
		log.Fatal(err)
	}
}
