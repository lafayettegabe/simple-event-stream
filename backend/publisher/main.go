package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lafayettegabe/simple-event-stream/utils"
)

const subject string = "updates"

func main() {
	nc := utils.ConnnectNATS()
	defer nc.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			id := uuid.New().String()
			err := nc.Publish(subject, []byte(id))
			if err != nil {
				log.Printf("Failed to publish message: %v", err)
				continue
			}
			fmt.Printf("Published UUID %s to subject %s\n", id, subject)
		}
	}
}
