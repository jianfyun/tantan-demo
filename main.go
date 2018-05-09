package main

import (
	"flag"
	"tantan-demo/api"
	"tantan-demo/config"
	"tantan-demo/lib/log"
	"tantan-demo/storage"
)

func main() {
	// Parse flags.
	var c string
	flag.StringVar(&c, "c", "", "config file path")
	flag.Parse()
	if c == "" {
		flag.Usage()
		return
	}
	// Load config file.
	if err := config.Load(c); err != nil {
		return
	}
	// Init logger.
	log.Init("demo")
	log.Info("tantan-demo starts")
	// Connect database
	storage.Connect()
	defer storage.Close()
	// Start HTTP Server and handle requests.
	api.Handle()
	log.Info("tantan-demo stops")
}
