package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	// 文字列の標準入力を3つ受け取る
	fmt.Scanf("%s", &S)

	// 入力値をスライスに分割する処理
	words := strings.Split(S, "")

	// 条件を確認する処理
	if words[0] == "R" {
		// Rが一番左
		fmt.Println("Yes")
	} else if words[1] == "R" && words[2] == "M" {
		// Rが真ん中で、味噌汁が右端
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
