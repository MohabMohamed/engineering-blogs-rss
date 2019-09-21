package main

import "fmt"

func main() {

	BlogsMap := ParseMainRepo()

	for _, blog := range BlogsMap.BlogsInfo {
		fmt.Println(blog.FeedUrl)
	}
}
