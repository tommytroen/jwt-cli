package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := New().Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
