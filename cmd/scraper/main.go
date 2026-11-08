package main

import (
	"fmt"
	"log"

	"github.com/E3TU/Kick-Live-Scraper/internal/browser"
)

func main() {
	b, err := browser.New()
	if err != nil {
		log.Fatalf("failed to launch browser: %v", err)
	}
	defer b.Close()

	page := b.MustPage("https://kick.com")
	info, err := page.Info()
	if err != nil {
		log.Fatalf("failed to get page info: %v", err)
	}

	fmt.Println(info.Title)
}
