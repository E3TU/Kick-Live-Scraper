package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://example.com")

	fmt.Println(page.MustInfo().Title)

}
