package main

import (
	"fmt"
)

func main() {
	var a int

	//標準入力を受け取る
	fmt.Scanf("%d", &a)

	count := strings.Split(a, "")

	//abcに入っている文字が1か0かを判断する
	//1だったら1を加算する処理を書く
	for _, str := range a {
		if str == "1" {
			count++
		}

	}
	fmt.Println(count)
}
