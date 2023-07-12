package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	configFile  string
	versionFlag bool
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.StringVar(&cmd.configFile, "c", "", "config file")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.Parse()
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}
