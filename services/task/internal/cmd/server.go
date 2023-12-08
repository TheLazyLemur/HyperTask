package main

import (
	"log"
	"log/slog"

	"github.com/thelazylemur/hypertask/services/task/internal/server"
)

func main() {
	slog.Info("Starting server on port :8081")

	s := server.NewServer(":8081")
	defer s.Stop()

	log.Fatal(s.Start())
}
