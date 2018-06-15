package main

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
	"time"
	tcclient "github.com/taskcluster/taskcluster-client-go"
	"github.com/taskcluster/taskcluster-client-go/tcqueue"
)

func main() {
	provisionerId := os.Args[1]
	workerType := os.Args[2]
	workerGroup := os.Args[3]
	workerId := os.Args[4]
	duration, err := time.ParseDuration(os.Args[5])
	if err != nil {
		log.Fatalf("Could not interpret duration '%v': %v", os.Args[5], err)
	}
	queue := tcqueue.NewFromEnv()
	quarantineDeadline := time.Now().Add(duration)
	qwr := &tcqueue.QuarantineWorkerRequest{
		QuarantineUntil: tcclient.Time(quarantineDeadline),
	}
	response, err := queue.QuarantineWorker(provisionerId, workerType, workerGroup, workerId, qwr)
	if err != nil {
		log.Fatalf("Could not quarantine worker due to: %v", err)
	}
	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Could not quarantine worker due to: %v", err)
	}
	fmt.Printf("Response: %s\n", bytes)
	fmt.Printf("Successfully updated quarantine for worker %v/%v/%v/%v\n", provisionerId, workerType, workerGroup, workerId)
}
