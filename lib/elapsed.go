package lib

import (
	"fmt"
	"time"
)

// ElapsedMessage は startTime からの経過時間を日本語の文字列で返します。
func ElapsedMessage(startTime time.Time) string {
	d := time.Since(startTime).Truncate(time.Second)

	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	switch {
	case h > 0:
		return fmt.Sprintf("会話開始から%d時間%d分経ってるよ", h, m)
	case m > 0:
		return fmt.Sprintf("会話開始から%d分%d秒経ってるよ", m, s)
	default:
		return fmt.Sprintf("会話開始から%d秒経ってるよ", s)
	}
}
