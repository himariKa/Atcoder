//文字列Sを標準入力で受け取る
//Tが空文字列であることを明示
//Tにいずれかを追加する処理
//SとTが同じである処理

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	// 文字列Sを標準入力で受け取る
	fmt.Scan(&str)

	words := []string{"dream", "dreamer", "erase", "eraser"}

	for len(str) > 0 {
		found := false
		for _, word := range words {
			if strings.HasSuffix(str, word) {
				str = str[:len(str)-len(word)]
				found = true
				break
			}

		}
		if !found {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
