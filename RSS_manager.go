package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		log.Fatalf("GATOR -- Error creating request\n%v", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("GATOR -- Error elaborating request\n%v", err)
	}
	res.Header.Set("User-Agent", "gator")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("GATOR -- Error extracting response body\n%v", err)
	}
	feed := RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		log.Fatalf("GATOR -- Error parsing response\n%v", err)
	}
}

func (f *RSSFeed) rssUnescape() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for item := range f.Channel.Item {

	}
}
