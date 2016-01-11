package main

import "github.com/PuerkitoBio/goquery"
import "strings"
import "net/url"
import "fmt"

func query(word string) {
	// for now, only Bing
	query, err := url.ParseQuery(word)
	if err != nil {
		fmt.Printf("invalid word: %s, error: %v", word, err)
	}

	uri := strings.Replace(config.dicts[0].url, "{word}", query.Encode(), 1)
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		fmt.Printf("query error: %v", err)
	}

	doc.Find("span.def").Each(func(i int, s *goquery.Selection) {
		text := s.Find("a, span").Text()
		fmt.Println(text)
	})
}