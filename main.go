package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/docker/machine/utils"
	"github.com/ehazlett/cert-tool/version"
)

func fatal(msg interface{}) {
	fmt.Printf("error: %v\n", msg)
	os.Exit(1)
}

func cmdGenerate(c *cli.Context) {
	outputDir := c.GlobalString("output-directory")
	if outputDir == "" {
		cli.ShowAppHelp(c)
		fatal("you must specify an output directory")
	}

	caCert := filepath.Join(outputDir, "ca.pem")
	caKey := filepath.Join(outputDir, "ca-key.pem")
	clientCert := filepath.Join(outputDir, "client.pem")
	clientKey := filepath.Join(outputDir, "client-key.pem")
	serverCert := filepath.Join(outputDir, "server.pem")
	serverKey := filepath.Join(outputDir, "server-key.pem")
	org := c.GlobalString("tls-ca-org")
	bits := c.GlobalInt("tls-bit-size")

	// check if ca cert exist and error if so
	f, fErr := os.Stat(caCert)
	if f != nil {
		resp := ""
		fmt.Printf("overwrite existing certs? (y/n): ")
		fmt.Scanln(&resp)
		if strings.ToLower(resp) != "y" {
			return
		}
	} else {
		// create output dir
		if err := os.MkdirAll(outputDir, 0700); err != nil {
			fatal(err.Error())
		}

	}

	if fErr != nil && !os.IsNotExist(fErr) {
		fatal(fErr)
	}

	// generate CA
	println("generating ca certificate/key")
	if err := utils.GenerateCACertificate(caCert, caKey, org, bits); err != nil {
		fatal(err)
	}

	// generate client cert
	println("generating client certificate/key")
	if err := utils.GenerateCert([]string{}, clientCert, clientKey, caCert, caKey, org, bits); err != nil {
		fatal(err)
	}

	// generate server cert if requested
	serverHosts := c.GlobalStringSlice("server")
	if len(serverHosts) > 0 {
		println("generating server certificate/key")
		if err := utils.GenerateCert(serverHosts, serverCert, serverKey, caCert, caKey, org, bits); err != nil {
			fatal(err)
		}
	}

	fmt.Printf("certificates successfully generated in %s\n", outputDir)
}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "Docker Cert Tool"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Author = "@ehazlett"
	app.Email = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output-directory, d",
			Value: "",
			Usage: "output directory for certs",
		},
		cli.StringFlag{
			Name:  "tls-ca-org, o",
			Value: "unknown",
			Usage: "CA organization for certs",
		},
		cli.IntFlag{
			Name:  "tls-bit-size, b",
			Value: 2048,
			Usage: "number of bits in the key (default: 2048)",
		},
		cli.StringSliceFlag{
			Name:  "server, s",
			Value: &cli.StringSlice{},
			Usage: "server host/ip for cert",
		},
	}

	app.Action = cmdGenerate
	if err := app.Run(os.Args); err != nil {
		fatal(err)
	}
}
