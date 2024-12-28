package main

import (
	"fmt"
)

func main() {
	var R, G, B int
	//標準入力を受け取る
	fmt.Scanf("%d %d %d", &R, &G, &B)

	var C string
	//標準入力を受け取る
	fmt.Scan(&C)

	if C == "Red" {
		if G > B {
			fmt.Println(B)
		} else {
			fmt.Println(G)
		}
	} else if C == "Green" {
		if R > B {
			fmt.Println(B)
		} else {
			fmt.Println(R)
		}
	} else {
		if R > G {
			fmt.Println(G)
		} else {
			fmt.Println(R)
		}
	}
}
