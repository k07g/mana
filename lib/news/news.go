package news

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

var rssURL = "https://www.nhk.or.jp/rss/news/cat0.xml"

type rssResponse struct {
	Channel struct {
		Items []struct {
			Title string `xml:"title"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetch(ctx context.Context) (*rssResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rssURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data rssResponse
	if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Today は NHK ニュースの最新トップ headlines を返します。
func Today(ctx context.Context) (string, error) {
	data, err := fetch(ctx)
	if err != nil {
		return "", err
	}

	items := data.Channel.Items
	if len(items) == 0 {
		return "", fmt.Errorf("ニュースを取得できませんでした")
	}

	const maxItems = 3
	if len(items) > maxItems {
		items = items[:maxItems]
	}

	var sb strings.Builder
	sb.WriteString("今日のニュースだよ！\n")
	for i, item := range items {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, item.Title))
	}
	return strings.TrimRight(sb.String(), "\n"), nil
}
