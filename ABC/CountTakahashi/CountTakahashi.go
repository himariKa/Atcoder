package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scan(&N)

	S := make([]string, N)
	//文字列の標準入力
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &S[i])
	}

	//Takahashiと等しい時にインクリメント
	count := 0
	for _, value := range S {
		if value == "Takahashi" {
			count++
		}
	}
	fmt.Println(count)
}
