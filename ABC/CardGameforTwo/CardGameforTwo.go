package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	//カードの数を標準入力で受け取る処理
	fmt.Scanf("%d", &N)

	a := make([]int, N)
	// スライスを使ってカードに書かれた数を標準入力から受け取る処理
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	//大きいもの順にaのスライスを並べ替える処理
	sort.Slice(a, func(i, j int) bool {
		//最大値を見つける処理
		return a[i] > a[j]
	})

	//アリスとボブの点数の合計点を求める処理
	var A, B int
	for i, value := range a {
		if i%2 == 0 {
			A += value
		} else  {
			B += value
		}
	}
	// アリスの手札の合計からボブの手札の合計を引く
	// 合計を引いた結果の数値を出力する処理
	fmt.Println(A - B)
}
