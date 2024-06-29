package main

import (
	"fmt"
)
	


func main()  {
	var a,b,c int
	var s string
	
	//標準入力を受け取る
	fmt.Scanf("%d",&a)
	fmt.Scanf("%d %d", &b ,&c)
	fmt.Scanf("%v",&s)

	fmt.Printf("%d %v\n",a+b+c, s)
}