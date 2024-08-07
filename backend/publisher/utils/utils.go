package utils

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func ConnnectNATS() *nats.Conn {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	var nc *nats.Conn
	var err error

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(natsURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to NATS, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to NATS after 5 attempts: %v", err)
	}

	log.Println("Connected to NATS server")

	return nc
}
