package main

import "log"

func run(port string, l *log.Logger) error {
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
	l := log.New(std.Stdout, "", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	
	if err := run(port, l); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}