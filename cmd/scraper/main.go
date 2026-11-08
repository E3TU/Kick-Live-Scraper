package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/E3TU/Kick-Live-Scraper/internal/scraper"
)

const (
	channel  = "test"
	interval = 60 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	runScrape(channel)

	for {
		select {
		case <-ticker.C:
			runScrape(channel)
		case <-ctx.Done():
			return
		}
	}
}

func runScrape(channel string) {
	result, err := scraper.Scrape(channel)
	if err != nil {
		log.Printf("scrape failed: %v", err)
		return
	}

	out, err := json.Marshal(result)
	if err != nil {
		log.Printf("marshal failed: %v", err)
		return
	}

	fmt.Println(string(out))
}
