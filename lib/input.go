package lib

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/k07g/mana/lib/greet"
)

func Input(ctx context.Context) error {
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
