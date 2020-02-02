---
title: REST API
layout: page
order: 20
---

# REST API

Use the [`net/http`](https://golang.org/pkg/net/http/) standard library to create external http calls. You can combine this with the [parsing of JSON](json.html) to work with external APIs.

The following example pulls in data from the Reddit API to get latest posts from golang sub. The key to understanding JSON decoding is understanding the data structure you receive the data. Once data structs are mapped, it becomes relatively straight forward

```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// create data structures to hold JSON data
// this needs to match the fields in the JSON feed
// see sample data fetch at bottom of this file

// create an individual entry data structure
// this does not need to hold every field, just the ones we want
type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

// the feed is the full JSON data structure
// this sets up the array of Entry types (defined above)
type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	// url of REST endpoint we are grabbing data from
	url := "https://www.reddit.com/r/golang/new.json"

	// fetch url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}
	// defer response close
	defer resp.Body.Close()

	// confirm we received an OK status
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	// read the entire body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	// create an empty instance of Feed struct
	// this is what gets filled in when unmarshaling JSON
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	// loop through the children and create entry objects
	for _, ed := range entries.Data.Children {
		entry := ed.Data
		log.Println(">>>")
		log.Println("Title   :", entry.Title)
		log.Println("Author  :", entry.Author)
		log.Println("URL     :", entry.URL)
		log.Printf("Comments: http://reddit.com%s \n", entry.Permalink)
	}
}
```

