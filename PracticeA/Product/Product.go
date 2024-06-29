package main

import (
	"fmt"
)
	


func main()  {
	var a,b int
	
	//標準入力を受け取る
	fmt.Scanf("%d %d",&a,&b)

	if a*b%2==0 {
		//積が偶数の時の処理
		fmt.Println("Even")
	}else{
		//奇数の時の処理
		fmt.Println("Odd")
	}
	
}