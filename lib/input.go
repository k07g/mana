package lib

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k07g/mana/lib/greet"
	"github.com/k07g/mana/lib/news"
	"github.com/k07g/mana/lib/weather"
)

func Input(ctx context.Context, version string, startTime time.Time) error {
	scanner := bufio.NewScanner(os.Stdin)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				return
			}
			input := scanner.Text()

			if input == "" {
				continue
			}

			if message, ok := greet.GreetHandler(input); ok {
				Say(message)
			}

			if strings.Contains(input, "バージョン") {
				Say("バージョン: " + version)
			}

			if strings.Contains(input, "会話時間") || strings.Contains(input, "どのくらい話した") || strings.Contains(input, "何分話した") {
				Say(ElapsedMessage(startTime))
			}

			if strings.Contains(input, "今日のニュース") || strings.Contains(input, "ニュース教えて") {
				msg, err := news.Today(ctx)
				if err != nil {
					Say("ニュースの取得に失敗しました: " + err.Error())
				} else {
					Say(msg)
				}
			}

			if strings.Contains(input, "明日の天気") {
				msg, err := weather.Tomorrow(ctx)
				if err != nil {
					Say("天気情報の取得に失敗しました: " + err.Error())
				} else {
					Say(msg)
				}
			} else if strings.Contains(input, "今日の天気") {
				msg, err := weather.Today(ctx)
				if err != nil {
					Say("天気情報の取得に失敗しました: " + err.Error())
				} else {
					Say(msg)
				}
			}

			if input == "quit" ||
				input == "exit" ||
				input == "おやすみ" ||
				input == "さようなら" ||
				input == "バイバイ" {
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		Say("\nお疲れさまー！おやすみ")
	case <-done:
	}
	return nil
}
