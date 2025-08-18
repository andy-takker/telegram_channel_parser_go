package domain

import "context"

type State interface {
	Get() int
	Set(int)
}

type Notifier interface {
	Send(ctx context.Context, text string) error
}

type Scraper interface {
	Fetch(ctx context.Context, channel ChannelUsername) ([]Post, error)
}
