package main

import (
	"flag"
	"log"

	local "github.com/bugfest/tor-controller/agents/onionbalance/local"
)

// onionbalance-manager main.
func main() {
	flag.Parse()

	//stopCh := signals.SetupSignalHandler()

	localManager := local.New()
	err := localManager.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
