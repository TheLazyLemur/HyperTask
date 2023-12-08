package main

import (
	"log"
	"log/slog"
)

func main() {
	slog.Info("Starting server on port :8081")

	s := NewServer(":8081")
	defer s.Stop()

	log.Fatal(s.Start())
}
