package greet

import "math/rand"

var goodMorningMessages = []string{
	"おはよー！",
	"おはようございます！",
	"グッドモーニング！",
	"今日もいい一日を！",
}

func goodMorning() string {
	return goodMorningMessages[rand.Intn(len(goodMorningMessages))]
}

var goodNightMessages = []string{
	"おやすみー！",
	"おやすみなさい！",
	"グッドナイト！",
	"ゆっくり休んでね！",
}

func goodNight() string {
	return goodNightMessages[rand.Intn(len(goodNightMessages))]
}

func goodBye() string {
	return "さようなら！またね！"
}
