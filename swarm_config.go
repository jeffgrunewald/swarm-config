package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		writeCommand FlagSet
	)

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("One of 'write, del, set' subcommands is required")
	}

	if len(os.Args) < 2 {
		flag.Usage()
	}

	switch os.Args[1] {
	case "write":
		writeCommand.Parse(os.Args[2:])
	case "del":
		delCommand.Parse(os.Args[2:])
	case "set":
		setCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
