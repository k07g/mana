package greet

// GreetHandler は、ユーザーの入力に応じて適切な挨拶メッセージを返します。
// 戻り値は、挨拶メッセージと、入力が挨拶に該当するかどうかを示すブール値です。
func GreetHandler(input string) (string, bool) {
	switch input {
	case "おはよ", "おはよう":
		return goodMorning(), true
	case "こんにちは":
		return "こんにちは！", true
	case "こんばんは":
		return "こんばんは！", true
	case "おやすみ":
		return goodNight(), true
	case "さようなら", "バイバイ":
		return goodBye(), true
	default:
		return "", false
	}
}
