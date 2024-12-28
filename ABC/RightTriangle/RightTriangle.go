package main

import (
	"fmt"
	"math"
)

func main() {
	var Xa, Ya, Xb, Yb, Xc, Yc float64

	fmt.Scanf("%f %f", &Xa, &Ya)
	fmt.Scanf("%f %f", &Xb, &Yb)
	fmt.Scanf("%f %f", &Xc, &Yc)

	// 3点を使って辺の長さを計算
	ab := math.Sqrt(math.Pow(Xb-Xa, 2) + math.Pow(Yb-Ya, 2))
	bc := math.Sqrt(math.Pow(Xc-Xb, 2) + math.Pow(Yc-Yb, 2))
	ca := math.Sqrt(math.Pow(Xa-Xc, 2) + math.Pow(Ya-Yc, 2))

	// 直角三角形かどうかを判定
	if isRightTriangle(ab, bc, ca) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 直角三角形かどうかを判定する関数
func isRightTriangle(a, b, c float64) bool {
	const epsilon = 1e-9 // 浮動小数点の比較のための許容誤差
	return math.Abs(a*a+b*b-c*c) < epsilon ||
		math.Abs(a*a+c*c-b*b) < epsilon ||
		math.Abs(b*b+c*c-a*a) < epsilon
}
