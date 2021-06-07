package main

import "fmt"

func main()  {
	var a int //默认是0
	fmt.Print("a = ", a)
	fmt.Printf(" type of a = %T\n", a)
	var b int = 100 //初始化
	fmt.Print("b = ", b)
	fmt.Printf(" type of b = %T\n", b)
	var c  = 0.01 //自动匹配类型初始化
	fmt.Print("c = ", c)
	fmt.Printf(" type of c = %T\n", c)
	d := "123" //省去var，自动匹配
	fmt.Print("d = ", d)
	fmt.Printf(" type of d = %T\n", d)
	/* ---多声明--- */
	var aa,bb int = -1, 0
	fmt.Printf(" type of aa = %T\n", aa)
	fmt.Printf(" type of bb = %T\n", bb)
	
}