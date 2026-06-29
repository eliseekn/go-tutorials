package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	u := launcher.New().Headless(false).Bin("/Applications/Brave Browser.app/Contents/MacOS/Brave Browser").MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://127.0.0.1:8001")

	page.MustEval(`() => document.querySelector('input[name="buttoncreatecrl1"]').click()`)
}
