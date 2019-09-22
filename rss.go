package main

import (
	"github.com/mmcdole/gofeed"
)

var fp *gofeed.Parser = gofeed.NewParser()

func GetRssFeed(feedURL string) []*gofeed.Item {

	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		panic(err)
	}
	return feed.Items

}
