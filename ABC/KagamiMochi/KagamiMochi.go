// 直径が同じ餅同士は重ねられない

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int

	//鏡餅の数を標準入力で受け取る処理
	fmt.Scan(&N)

	d := make([]int, N)

	//鏡餅の直径を標準入力で受け取る処理
	for i := 0; i < N; i++ {
		fmt.Scan(&d[i])
	}

	//大きい数順に並び替え
	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	count := 1
	for i := 0; i < N-1; i++ {
		//[10,8,8,6]
		if d[i] > d[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
