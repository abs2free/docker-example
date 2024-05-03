package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
		//APIKey: "Xq7aZsAiRoO-i2YorCI7YQ",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// API Key should have cluster monitoring rights
	infores, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fmt.Println(infores)
}
