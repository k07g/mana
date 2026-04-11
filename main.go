package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/k07g/mana/lib"
)

func main() {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lib.Say("おはよー")

	if err := lib.Input(ctx); err != nil {
		fmt.Println("入力処理に失敗しました:", err)
	}
}
