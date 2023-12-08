package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Starting server on port %s\n", ":8081")

	s := NewServer(":8081")
	defer s.Stop()

	log.Fatal(s.Start())
}
