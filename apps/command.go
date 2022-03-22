package main

import "flag"

type CmdParams struct {
	port *string
}

func (cmd *CmdParams) Parse() bool {
	cmd.port = flag.String("port", "8080", "Server port")
	// Add additional arguments

	flag.Parse()

	return flag.NArg() == 0
}

func (cmd *CmdParams) PrintDefaults() {
	flag.PrintDefaults()
}
