package main

import (
	"log"
	"net/http"
	"time"

	"temporal-ip-geolocation/iplocate"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the Temporal client
	c, err := client.Dial(client.Options{
		HostPort:          "localhost:7233",
		ConnectionOptions: client.ConnectionOptions{GetSystemInfoTimeout: 30 * time.Second},
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create the Temporal worker
	w := worker.New(c, iplocate.TaskQueueName, worker.Options{})

	// inject HTTP client into the Activities Struct
	activities := &iplocate.IPActivities{
		HTTPClient: http.DefaultClient,
	}

	// Register Workflow and Activities
	w.RegisterWorkflow(iplocate.GetAddressFromIP)
	w.RegisterActivity(activities)

	// Start the Worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Temporal worker", err)
	}
}
