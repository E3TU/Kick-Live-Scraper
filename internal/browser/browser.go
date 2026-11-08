package browser

import (
	"github.com/go-rod/rod"
)

type Browser struct {
	browser *rod.Browser
}

func New() (*Browser, error) {
	b := rod.New()
	if err := b.Connect(); err != nil {
		return nil, err
	}
	return &Browser{browser: b}, nil
}

func (b *Browser) MustPage(url string) *rod.Page {
	return b.browser.MustPage(url)
}

func (b *Browser) Close() error {
	return b.browser.Close()
}
