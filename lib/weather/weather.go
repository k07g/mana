package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var apiURL = "https://wttr.in/?format=j1&lang=ja"

type langJa struct {
	Value string `json:"value"`
}

type response struct {
	CurrentCondition []struct {
		TempC  string   `json:"temp_C"`
		LangJa []langJa `json:"lang_ja"`
	} `json:"current_condition"`
	Weather []struct {
		MaxTempC string `json:"maxtempC"`
		MinTempC string `json:"mintempC"`
		Hourly   []struct {
			LangJa []langJa `json:"lang_ja"`
		} `json:"hourly"`
	} `json:"weather"`
}

func fetch(ctx context.Context) (*response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func Today(ctx context.Context) (string, error) {
	data, err := fetch(ctx)
	if err != nil {
		return "", err
	}

	if len(data.CurrentCondition) == 0 || len(data.CurrentCondition[0].LangJa) == 0 {
		return "", fmt.Errorf("天気情報を取得できませんでした")
	}

	c := data.CurrentCondition[0]
	return fmt.Sprintf("今日の天気は %s、気温は %s°Cだよ", c.LangJa[0].Value, c.TempC), nil
}

func Tomorrow(ctx context.Context) (string, error) {
	data, err := fetch(ctx)
	if err != nil {
		return "", err
	}

	if len(data.Weather) < 2 {
		return "", fmt.Errorf("明日の天気情報を取得できませんでした")
	}

	w := data.Weather[1]
	// 正午(index=4)の天気説明を使用する
	if len(w.Hourly) < 5 || len(w.Hourly[4].LangJa) == 0 {
		return "", fmt.Errorf("明日の天気情報を取得できませんでした")
	}

	desc := w.Hourly[4].LangJa[0].Value
	return fmt.Sprintf("明日の天気は %s、最高 %s°C / 最低 %s°Cだよ", desc, w.MaxTempC, w.MinTempC), nil
}
