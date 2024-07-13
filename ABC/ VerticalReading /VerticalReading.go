package main

import (
	"fmt"
)

func main() {
	var S, T string
	// 文字列SとTを標準入力する
	fmt.Scanf("%s %s", &S, &T)

	// SとTの文字数を調べる処理
	lenS := len(S)
	lenT := len(T)

	// Tの文字数から1文字引いた数を調べる処理
	wordt := lenT - 1

	// Sの文字数を上の処理で調べた数まで分解する処理
	var concatenated string
	if lenS >= wordt {
		concatenated = S[lenS-wordt:]
	}

	// 最後の文字を連結したものとTを比較して同じだったらYesを出力する
	if concatenated == T {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
