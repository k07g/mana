package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	verbose := flag.Bool("verbose", false, "詳細出力を有効にする")
	flag.Parse()

	fmt.Println("=== Go対話式計算機 ===")
	fmt.Println("使用方法: <数値1> <演算子> <数値2>")
	fmt.Println("演算子: + - * / %")
	fmt.Println("'quit' または 'exit' で終了します")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if input == "quit" || input == "exit" {
			fmt.Println("終了します。さようなら！")
			break
		}

		result, err := Calculate(input, *verbose)
		if err != nil {
			fmt.Printf("エラー: %v\n", err)
			continue
		}

		fmt.Printf("結果: %v\n", result)
	}
}

func Calculate(input string, verbose bool) (float64, error) {
	parts := strings.Fields(input)

	if len(parts) != 3 {
		return 0, fmt.Errorf("入力形式が正しくありません。<数値1> <演算子> <数値2> の形で入力してください")
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("最初の値が数値ではありません: %s", parts[0])
	}

	operator := parts[1]

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("2番目の値が数値ではありません: %s", parts[2])
	}

	if verbose {
		fmt.Printf("入力: %f %s %f\n", num1, operator, num2)
	}

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("ゼロで除算することはできません")
		}
		result = num1 / num2
	case "%":
		if num2 == 0 {
			return 0, fmt.Errorf("ゼロで除算することはできません")
		}
		result = float64(int64(num1) % int64(num2))
	default:
		return 0, fmt.Errorf("サポートされない演算子: %s", operator)
	}

	return result, nil
}
