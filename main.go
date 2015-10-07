package main

// Main entry point for the app. Handles command-line options and starts the web listener

import (
	"flag"
	"fmt"
	"os"
)

var (
	httpPort    int
	botUsername string
	botIcon     string
)

func main() {
	// Parse command-line options
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: ./slack-leto -port=8002\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&httpPort, "port", 8002, "The HTTP port on which to listen")
	flag.StringVar(&botUsername, "botUsername", "leto", "The name of the bot when it speaks")
	flag.StringVar(&botIcon, "botIcon", "https://example.com/icon.png", "The web-accessible URL of the icon to use for the bot")

	flag.Parse()

	if httpPort == 0 {
		flag.Usage()
		os.Exit(2)
	}

	// Start the webserver
	StartServer(httpPort)
}
