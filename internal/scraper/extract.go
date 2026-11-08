package scraper

import (
	"time"

	"github.com/go-rod/rod"
)

const liveSelector = "#injected-embedded-channel-player-video"

func ExtractLiveStatus(page *rod.Page) (bool, error) {
	_, err := page.Timeout(5 * time.Second).Element(liveSelector)
	if err != nil {
		return false, nil
	}
	return true, nil
}
