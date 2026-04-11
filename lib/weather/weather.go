package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	CurrentCondition []struct {
		TempC  string `json:"temp_C"`
		LangJa []struct {
			Value string `json:"value"`
		} `json:"lang_ja"`
	} `json:"current_condition"`
}

func Today(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://wttr.in/?format=j1&lang=ja", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if len(data.CurrentCondition) == 0 || len(data.CurrentCondition[0].LangJa) == 0 {
		return "", fmt.Errorf("天気情報を取得できませんでした")
	}

	c := data.CurrentCondition[0]
	return fmt.Sprintf("今日の天気は %s、気温は %s°Cだよ", c.LangJa[0].Value, c.TempC), nil
}
