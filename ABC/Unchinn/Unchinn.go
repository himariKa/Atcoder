package main

import "fmt"

func main() {
	var N int
	var M int
	var S int

	fmt.Scan(&N, &M)

	//路線の数
	a := make([]int, N)

	//駅の数
	// b := make([]int, M)

	//路線の数まで入力値を受け付ける
	for i := 0; i <= N; i++ {
		//駅の数まで入力値を受け付ける
		for j := 0; j <= M; j++ {
			fmt.Scan(&S)
			a = append(a, S)
			fmt.Print(a)
		}
	}
}
