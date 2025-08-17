package scraper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	MsgID int
	URL   string
	Text  string
}

type Scraper struct {
	Client *http.Client
}

func New(c *http.Client) *Scraper {
	return &Scraper{Client: c}
}

func (s *Scraper) Fetch(ctx context.Context, channel string) ([]Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://t.me/s/"+channel, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; telegram-channel-poller/1.0)")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return ParseDocument(doc, channel), nil
}
