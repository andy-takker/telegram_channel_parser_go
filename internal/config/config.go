package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	BotToken        string
	ChatID          string
	ChannelUsername string
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
		BotToken:        getStrict("APP_BOT_TOKEN"),
		ChatID:          getStrict("APP_CHAT_ID"),
		ChannelUsername: strings.TrimPrefix(getStrict("APP_CHANNEL_USERNAME"), "@"),
		PollInterval:    time.Duration(getInt("APP_POLL_SECONDS", 5)) * time.Second,
		StorePath:       strings.TrimSpace(os.Getenv("APP_STATE_FILE")),
		Keyword:         getStrict("APP_KEYWORD"),
	}
}
