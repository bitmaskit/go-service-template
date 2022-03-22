package main

import (
	"fmt"
	"log"
	"os"
)

func run(port string, l *log.Logger) error {
	l.Println("Running on port: ", port)
	// ...

	return nil
}

func main() {
	// Command Line Args
	var cmd CmdParams
	if !cmd.Parse() {
		cmd.PrintDefaults()
		return
	}

	// Logger
	l := log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)

	if err := run(*cmd.port, l); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
