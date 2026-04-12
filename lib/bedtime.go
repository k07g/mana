package lib

// BedtimeMessage は hour が就寝を促す時間帯（22時〜翌4時）であれば
// メッセージと true を返す。それ以外は空文字と false を返す。
func BedtimeMessage(hour int) (string, bool) {
	if hour >= 22 || hour < 4 {
		return "もうこんな時間！早く寝ないと体に悪いよ", true
	}
	return "", false
}
