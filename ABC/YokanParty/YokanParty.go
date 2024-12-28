package main

import "fmt"

func main() {
	var L int
	var N int
	var K int
	A := make([]int, N)
	//切れ目の数を入力から受け取る
	fmt.Scan(&N)
	//羊羹の長さを入力から受け取る
	fmt.Scan(&L)
	//えらぶ羊羹の数
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}
