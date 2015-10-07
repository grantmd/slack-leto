package main

// Main entry point for the app. Handles command-line options and starts the web listener

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	httpPort    int
	botUsername string
	botIcon     string
	responded   map[string]int
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
}

func main() {
	// Parse command-line options
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: ./slack-leto -port=8002\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&httpPort, "port", 8002, "The HTTP port on which to listen")
	flag.StringVar(&botUsername, "botUsername", "leto", "The name of the bot when it speaks")
	flag.StringVar(&botIcon, "botIcon", "", "The web-accessible URL of the icon to use for the bot")

	flag.Parse()

	if httpPort == 0 {
		flag.Usage()
		os.Exit(2)
	}

	responded = make(map[string]int)

	// Start the webserver
	StartServer(httpPort)
}
