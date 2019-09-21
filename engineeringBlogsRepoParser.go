package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Struct to reprsent the single blog
//info in the xml that contains the blog links
type BlogInfo struct {
	Title   string `xml:"title,attr"`
	Url     string `xml:"htmlUrl,attr"`
	FeedUrl string `xml:"xmlUrl,attr"`
}

//Struct to reprsent the outter structure of the blogs
//info in the xml that contains the blog links
type EngineeringBlogsMap struct {
	XMLName   xml.Name   `xml:"opml"`
	BlogsInfo []BlogInfo `xml:"body>outline>outline"`
}

// Parse https://github.com/MohabMohamed/engineering-blogs
//repo to get all the RSS urls for the blogs in it
func ParseMainRepo() EngineeringBlogsMap {
	resp, err := http.Get("https://raw.githubusercontent.com/MohabMohamed/engineering-blogs/master/engineering_blogs.opml")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var BlogsMap EngineeringBlogsMap
	err = xml.Unmarshal(bytes, &BlogsMap)
	if err != nil {
		panic(err)
	}
	return BlogsMap
}
