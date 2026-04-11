package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/k07g/mana/lib"
)

var version = "v0.1.1"

func main() {
	showVersion := flag.Bool("version", false, "バージョンを表示する")
	flag.BoolVar(showVersion, "v", false, "バージョンを表示する")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lib.Say("おはよー")

	if err := lib.Input(ctx, version); err != nil {
		fmt.Println("入力処理に失敗しました:", err)
	}
}
