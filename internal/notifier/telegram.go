package notifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Telegram struct {
	Client         *http.Client
	Token          string
	ChatID         string
	DisablePreview bool
}

type sendMessageReq struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

func NewTelegram(c *http.Client, token, chatID string) *Telegram {
	return &Telegram{Client: c, Token: token, ChatID: chatID, DisablePreview: true}
}

func (t *Telegram) Send(ctx context.Context, text string) error {
	body := sendMessageReq{
		ChatID:                t.ChatID,
		Text:                  text,
		DisableWebPagePreview: t.DisablePreview,
	}
	j, _ := json.Marshal(body)
	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", url.PathEscape(t.Token))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("sendMessage failed: %s", resp.Status)
	}
	return nil
}

func ComposeMessage(msgID int, url, text string) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "#%d:\n%s\n\n", msgID, url)
	if text != "" {
		b.WriteString(text)
	}
	return b.String()
}
