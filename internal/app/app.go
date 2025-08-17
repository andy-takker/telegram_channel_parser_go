package app

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/andy-takker/telegram_channel_parser_go/internal/notifier"
	"github.com/andy-takker/telegram_channel_parser_go/internal/scraper"
)

type State interface {
	Get() int
	Set(int)
}

type Notifier interface {
	Send(ctx context.Context, text string) error
}

type App struct {
	Scraper  *scraper.Scraper
	Notifier Notifier
	Channel  string
	Every    time.Duration
	State    State
	Keyword  string
}

func (a *App) Tick(ctx context.Context) error {
	posts, err := a.Scraper.Fetch(ctx, a.Channel)
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		return errors.New("no posts found")
	}
	last := a.State.Get()
	var news []scraper.Post
	for _, p := range posts {
		if p.MsgID > last {
			// === фильтрация по ключевому слову ===
			if a.Keyword != "" {
				if !containsKeyword(p.Text, a.Keyword) {
					continue // пропускаем пост
				}
			}
			news = append(news, p)
		}
	}
	if len(news) == 0 {
		log.Printf("Нет новых подходящих постов (last_seen=%d)", last)
		return nil
	}
	for _, p := range news {
		msg := notifier.ComposeMessage(p.MsgID, p.URL, p.Text)
		if err := a.Notifier.Send(ctx, msg); err != nil {
			return err
		}
		log.Printf("→ Отправлено: %s", p.URL)
	}
	a.State.Set(news[len(news)-1].MsgID)
	return nil
}

func (a *App) Run(ctx context.Context) {
	ticker := time.NewTicker(a.Every)
	defer ticker.Stop()

	_ = a.Tick(ctx) // первый проход
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := a.Tick(ctx); err != nil {
				log.Printf("Ошибка тикера: %v", err)
			}
		}
	}
}

func containsKeyword(text, keyword string) bool {
	return strings.Contains(
		strings.ToLower(text),
		strings.ToLower(keyword),
	)
}
