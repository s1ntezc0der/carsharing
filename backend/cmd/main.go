package main

import (
	"log"

	"github.com/s1ntezc0der/car-rental/backend/run"
)

func main() {
	if err := run.Run(); err != nil {
		log.Fatal(err)
	}
}
