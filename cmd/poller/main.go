package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/andy-takker/telegram_channel_parser_go/internal/app"
	"github.com/andy-takker/telegram_channel_parser_go/internal/config"
	"github.com/andy-takker/telegram_channel_parser_go/internal/logx"
	"github.com/andy-takker/telegram_channel_parser_go/internal/notifier"
	"github.com/andy-takker/telegram_channel_parser_go/internal/scraper"
	"github.com/andy-takker/telegram_channel_parser_go/internal/state"
	"github.com/andy-takker/telegram_channel_parser_go/pkg/httputil"
)

func main() {
	logx.Init()
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	client := httputil.NewDefaultClient()
	s := scraper.New(client)
	n := notifier.NewTelegram(client, cfg.BotToken, cfg.ChatID)

	var st app.State
	if cfg.StorePath != "" {
		fst, err := state.NewFileStore(cfg.StorePath)
		if err != nil {
			// если файл не удаётся, откатимся на память
			st = state.NewMemory()
		} else {
			st = fst
		}
	} else {
		st = state.NewMemory()
	}

	app := &app.App{
		Scraper:  s,
		Notifier: n,
		Channel:  cfg.ChannelUsername,
		Every:    cfg.PollInterval,
		State:    st,
		Keyword:  cfg.Keyword,
	}
	app.Run(ctx)
}
