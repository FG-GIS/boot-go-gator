package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
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
	feed := &RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return feed, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return feed, err
	}
	res.Header.Set("User-Agent", "gator")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return feed, err
	}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return feed, err
	}
	feed.rssUnescape()
	return feed, nil
}

func (f *RSSFeed) rssUnescape() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for idx, item := range f.Channel.Item {
		f.Channel.Item[idx].Title = html.UnescapeString(item.Title)
		f.Channel.Item[idx].Description = html.UnescapeString(item.Description)
	}
}
