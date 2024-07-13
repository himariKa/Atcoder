package main

import (
	"fmt"
)

func main() {
	//入っていたお札N枚を標準入力から受け取る処理
	var N int
	fmt.Scan(&N)
	//入っていたお札の合計Y円を標準入力から受け取る処理
	var Y int
	fmt.Scan(&Y)

	// 一万円札の数の最大値を調べて最小値まで行ったら五千円札の最大値を調べて残りは千円札にする処理
	// 千円札の数が0枚以上で合計値がYと同じの時、そのそれぞれのお札の枚数を出力する処理
	for i := Y / 10000; i >= 0; i-- {
		for j := N - i; j >= 0; j-- {
			k := N - i - j
			if k >= 0 && 10000*i+5000*j+1000*k == Y {
				fmt.Printf("%d %d %d", i, j, k)
				return
			}

		}
	}
	fmt.Printf("%d %d %d", -1, -1, -1)
}
