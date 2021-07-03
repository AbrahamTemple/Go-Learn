package main

import "fmt"

/**
方法声明及使用
 */

//单返回值
func foo(a string,b int) int {
	fmt.Print("a = ",a)
	fmt.Print("\n")
	fmt.Print("b = ",b)
	fmt.Print("\n")
	c := b+6
	return c
}

//多返回值匿名
func foo2(a string,b int) (int,int) {
	fmt.Print("a = ",a)
	fmt.Print("\n")
	fmt.Print("b = ",b)
	fmt.Print("\n")
	c := b*6
	return b,c
}

//多返回值署名
func foo3(a string,b int) (r1 int,r2 int) {
	fmt.Print("a = ",a)
	fmt.Print("\n")
	fmt.Print("r1 = ",r1)
	fmt.Print("\n")
	fmt.Print("r2 = ",r2)
	fmt.Print("\n")

	r1 = b*6+6
	r2 = b*6-6

	return
}

//同类型合并返回值署名
func foo4(a string,b int) (r1 ,r2 int) {
	fmt.Print("a = ",a)
	fmt.Print("\n")
	fmt.Print("r1 = ",r1)
	fmt.Print("\n")
	fmt.Print("r2 = ",r2)
	fmt.Print("\n")

	r1 = b*8+8
	r2 = b*8-8

	return
}

func main() {
	c := foo("单返回值",2)
	fmt.Print("c1 = ",c)
	fmt.Print("\n")

	b, c := foo2("多返回值匿名", 2)
	fmt.Print("b2 = ",b)
	fmt.Print("\n")
	fmt.Print("c2 = ",c)
	fmt.Print("\n")

	b, c = foo3("多返回值署名", 2)
	fmt.Print("b3 = ",b)
	fmt.Print("\n")
	fmt.Print("c3 = ",c)
	fmt.Print("\n")

	b, c = foo4("同类型合并返回值署名", 2)
	fmt.Print("b3 = ",b)
	fmt.Print("\n")
	fmt.Print("c3 = ",c)
	fmt.Print("\n")
}
