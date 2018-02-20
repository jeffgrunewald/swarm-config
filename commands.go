package main

import (
	"flag"
	"fmt"
)

//var userHome, _ = getUserHome()

func init() {
	// define subcommands
	//	writeCommand := flag.NewFlagSet("write", flag.ExitOnError)
	delCommand := flag.NewFlagSet("del", flag.ExitOnError)
	setCommand := flag.NewFlagSet("set", flag.ExitOnError)

	// write subcommand flags
	//	writeHostPtr := writeCommand.String("endpoint", "127.0.0.1", "Docker endpoint address.")
	//	writePortPtr := writeCommand.Int("port", 2375, "Docker endpoint TCP port")
	//	writeTLSPtr := writeCommand.Bool("tls", false, "Encrypt communication with docker.")
	//	writeCertPtr := writeCommand.String("cert-path", userHome, "Path to TLS certificates.")

	// del subcommand flags
	delEndpointPtr := delCommand.String("endpoint", "", "Endpoint config to remove")

	// set subcommand flags
	setEndpointPtr := setCommand.String("endpoint", "", "Endpoint to set as current context")

	fmt.Print(delEndpointPtr, setEndpointPtr)
}

//func getUserHome() (string, error) {
//	usr, err := user.Current()
//	if err != nil {
//		log.Fatal(err)
//	}
//	return usr.HomeDir, nil
//}
