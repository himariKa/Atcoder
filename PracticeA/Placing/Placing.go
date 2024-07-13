// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var S, T string
// 	// 文字列SとTを標準入力する
// 	fmt.Scanf("%s %s", &S, &T)

// 	// SとTの文字数を調べる処理
// 	lenS := len(S)
// 	lenT := len(T)

// 	// Tの文字数から1文字引いた数を調べる処理
// 	wordt := lenT - 1

// 	// Sの文字数を上の処理で調べた数まで分解する処理
// 	var concatenated string
// 	if lenS >= wordt {
// 		concatenated = S[:lenS-wordt]
// 	}

// 	// 最後の文字を連結したものとTを比較して同じだったらYesを出力する
// 	if concatenated == T {
// 		fmt.Println("Yes")
// 	} else {
// 		fmt.Println("No")
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scanf("%s %s", &s, &t)

	// 文字の区切りごとに繰り返し
	for word := 1; word < len(s); word++ {
		// 長さと取得インデックスの数繰り返し
		for c := 0; c < word; c++ {

			n := ""
			for i := c; i < len(s); i += word {
				//最後の文字の連結をnに入れる
				n += string(s[i])
			}
			if n == t {
				fmt.Print("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
