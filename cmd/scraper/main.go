package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/E3TU/Kick-Live-Scraper/internal/scraper"
)

func main() {
	channel := "test"

	result, err := scraper.Scrape(channel)
	if err != nil {
		log.Fatalf("scrape failed: %v", err)
	}

	out, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("marshal failed: %v", err)
	}

	fmt.Println(string(out))
}
