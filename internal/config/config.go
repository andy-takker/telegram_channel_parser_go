package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andy-takker/telegram_channel_parser_go/internal/domain"
)

type Config struct {
	BotToken        domain.BotToken
	ChatID          domain.ChatID
	ChannelUsername domain.ChannelUsername
	PollInterval    time.Duration
	StorePath       string
	Keyword         string
}

func getStrict(key string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}

func getInt(key string, def int) int {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return def
	}
	return n
}

func Load() Config {
	return Config{
		BotToken:        domain.BotToken(getStrict("APP_BOT_TOKEN")),
		ChatID:          domain.ChatID(getStrict("APP_CHAT_ID")),
		ChannelUsername: domain.ChannelUsername(strings.TrimPrefix(getStrict("APP_CHANNEL_USERNAME"), "@")),
		PollInterval:    time.Duration(getInt("APP_POLL_SECONDS", 5)) * time.Second,
		StorePath:       strings.TrimSpace(os.Getenv("APP_STATE_FILE")),
		Keyword:         getStrict("APP_KEYWORD"),
	}
}
