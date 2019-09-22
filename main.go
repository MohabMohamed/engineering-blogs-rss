package main

import "github.com/mmcdole/gofeed"

func main() {

	BlogsMap := ParseMainRepo()
	BlogsRssMap := make(map[string][]*gofeed.Item)
	for _, blog := range BlogsMap.BlogsInfo {
		BlogsRssMap[blog.Title] = GetRssFeed(blog.FeedUrl)
	}
}
