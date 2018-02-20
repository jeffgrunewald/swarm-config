package main

import (
	"flag"
	"log"
	"os/user"
)

var userHome, _ = getUserHome()

func writeCmd() {
	writeCommand := flag.NewFlagSet("write", flag.ExitOnError)

	// write subcommand flags
	writeHostPtr := writeCommand.String("endpoint", "127.0.0.1", "Docker endpoint address.")
	writePortPtr := writeCommand.Int("port", 2375, "Docker endpoint TCP port")
	writeTLSPtr := writeCommand.Bool("tls", false, "Encrypt communication with docker.")
	writeCertPtr := writeCommand.String("cert-path", userHome, "Path to TLS certificates.")
}

func getUserHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir, nil
}
