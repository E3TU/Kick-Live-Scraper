package scraper

import (
	"github.com/E3TU/Kick-Live-Scraper/internal/browser"
	"github.com/E3TU/Kick-Live-Scraper/internal/models"
)

func Scrape(channel string) (models.Channel, error) {
	b, err := browser.New()
	if err != nil {
		return models.Channel{}, err
	}
	defer b.Close()

	page := b.MustPage("https://kick.com/" + channel)

	live, err := ExtractLiveStatus(page)
	if err != nil {
		return models.Channel{}, err
	}

	return models.Channel{Name: channel, Live: live}, nil
}
