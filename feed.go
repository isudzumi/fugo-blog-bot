package main

import (
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

const intervalMin = 1

var (
	rssUrl = os.Getenv("FUGO_BLOG_RSS_LINK")
)

func getTime() (time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(jst).Add(intervalMin * time.Minute * -1)
	return t, err
}

func CheckIfRSSUpdated() (bool, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssUrl)

	if err != nil {
		return false, err
	}

	t, err := getTime()

	if err != nil {
		return false, err
	}

	if feed.Items == nil {
		return false, nil
	}

	if feed.Items[0].PublishedParsed.Before(t) {
		return false, nil
	}

	return true, nil
}
