# Telegram Channerl Parser (Go)

Telegram Channel Parser is Go service that **parses a public Telegram channel** (via https://t.me/s/<channel>) and forwards new posts to a chat using the **Telegram Bot API**.

Features:

- 🔍 Polls the public web version of a Telegram channel (t.me/s/...)
- 🆕 Tracks only new messages (last_seen_id)
- 📝 Sends notifications to the configured chat via Bot API
- 🎯 Filters messages by keyword (APP_KEYWORD, case-insensitive)
- 💾 Stores state in memory or in a JSON file (supports persistence)
- ⚡ Lightweight, no need for official Telegram API access
- 🐳 Easy to run locally or in Docker

## ⚙️ Environment Variables

| Variable               | Required | Default | Description                                                                        |
| ---------------------- | -------- | ------- | ---------------------------------------------------------------------------------- |
| `APP_BOT_TOKEN`        | ✅        | -       | Telegram Bot token (123456:ABC...)                                                 |
| `APP_CHAT_ID`          | ✅        | -       | Chat/channel ID where messages will be sent                                        |
| `APP_CHANNEL_USERNAME` | ✅        | -       | Channel username without @ (e.g., mychannel)                                       |
| `APP_POLL_SECONDS`     | ❌        | 5       | Polling interval in seconds                                                        |
| `APP_STATE_FILE`       | ❌        | -       | Path to JSON file for storing last_seen_id (e.g. ./last_seen.json)                 |
| `APP_KEYWORD`          | ❌        | -       | Keyword for filtering posts (case-insensitive). If empty → all posts are forwarded |

## 🚀 Running

### Locally

```bash
# install dependencies
go mod tidy

# export environment variables
export APP_BOT_TOKEN="123456:ABC..."
export APP_CHAT_ID="123456789"
export APP_CHANNEL_USERNAME="mychannel"
export APP_POLL_SECONDS=5
export APP_KEYWORD="bitcoin"

# run
go run ./cmd/poller/main.go
```

### Docker

```bash
# build image
docker build -t telegram-channel-parser .

# run container
docker run -it --rm \
  -e APP_BOT_TOKEN="123456:ABC..." \
  -e APP_CHAT_ID="123456789" \
  -e APP_CHANNEL_USERNAME="mychannel" \
  -e APP_KEYWORD="bitcoin" \
  -v $(pwd)/data:/data \
  telegram-channel-parser
```

> When mounting /data, the service will persist last_seen.json with the last seen message ID.

## 🧰 Makefile Commands

- `make fmt` - format code (`go fmt ./...`)
- `make lint` - run linter (`golangci-lint run`)
- `make test` - run tests with `-race` and coverage
- `make build` - build binary into `./bin/poller`
- `make run` - run the service locally
- `make clean` - remove build artifacts

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

With race detector and coverage:

```bash
go test -race -cover ./...
```

## 📦 Project Structure

```
telegram_channel_parser_go/
├── cmd/poller/        # entrypoint
├── internal/
│   ├── app/           # business logic (polling loop, filtering)
│   ├── config/        # configuration
│   ├── scraper/       # parse t.me/s/<channel>
│   ├── notifier/      # Telegram Bot API client
│   ├── state/         # last_seen_id storage
│   └── logx/          # logging helpers
├── pkg/httputil/      # helper http.Client
├── deploy/            # Dockerfile, docker-compose.yml
├── scripts/           # dev scripts
├── Makefile
└── README.md
```
