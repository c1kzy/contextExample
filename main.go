package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"contextexample/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)

	if err := server.startServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)

	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}


