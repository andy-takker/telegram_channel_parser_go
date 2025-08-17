package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(doc *goquery.Document, channel string) []Post {
	var posts []Post
	doc.Find(".tgme_widget_message_wrap").Each(func(_ int, wrap *goquery.Selection) {
		root := wrap.Find(".tgme_widget_message")
		if root.Length() == 0 {
			return
		}
		dp, _ := root.Attr("data-post") // channel/12345
		parts := strings.SplitN(dp, "/", 2)
		if len(parts) != 2 {
			return
		}
		id, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return
		}
		var url string
		root.Find(".tgme_widget_message_date a").EachWithBreak(func(_ int, a *goquery.Selection) bool {
			if href, ok := a.Attr("href"); ok && strings.Contains(href, "/"+channel+"/") {
				url = href
				return false
			}
			return true
		})
		if url == "" {
			url = fmt.Sprintf("https://t.me/%s/%d", channel, id)
		}
		text := normalize(root.Find(".tgme_widget_message_text").Text())
		if text == "" {
			text = normalize(root.Text())
		}
		runes := []rune(text)
		if len(runes) > 600 {
			text = string(runes[:600]) + "…"
		}
		posts = append(posts, Post{MsgID: id, URL: url, Text: text})
	})
	// простая сортировка вставками по возрастанию id
	for i := 1; i < len(posts); i++ {
		for j := i; j > 0 && posts[j-1].MsgID > posts[j].MsgID; j-- {
			posts[j-1], posts[j] = posts[j], posts[j-1]
		}
	}
	return posts
}

func normalize(s string) string {
	var b strings.Builder
	space := false
	for _, r := range s {
		if r == '\n' || r == '\r' || r == '\t' || r == '\v' || r == '\f' || r == ' ' {
			if !space {
				b.WriteByte(' ')
				space = true
			}
		} else {
			b.WriteRune(r)
			space = false
		}
	}
	return strings.TrimSpace(b.String())
}
