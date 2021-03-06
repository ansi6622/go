package main

import (
	"flag"
	"fmt"
	"os"
)

// START OMIT
func main() {
	var (
		cmd  string = "website"
		port int    = 8000
	)

	fs := flag.NewFlagSet("default", flag.ExitOnError)
	fs.StringVar(&cmd, "cmd", cmd, "the command to run")
	fs.IntVar(&port, "p", port, "the port to run on")
	fs.Parse(os.Args[1:])

	fmt.Printf("Running %q on port %d", cmd, port)
}

// END OMIT
