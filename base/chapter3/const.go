package main

import "fmt"

/**
常量声明及使用
 */

const(
	//iota关键字以下的会累计加5，默认是0(BEGIN)
	BEGIN = 5*iota
	ONE
	TWO
)

const(
	//iota关键字以下的会累计加1，默认是0(BEGIN)
	BEGINS = 5+iota
	ONCE
	TWICE
)

func main(){
	const len int = 10 //常量只可读
	fmt.Print("length = ",len)
	fmt.Print("\n")
	fmt.Print("BEGIN = ",BEGIN)
	fmt.Print("\n")
	fmt.Print("ONE = ",ONE)
	fmt.Print("\n")
	fmt.Print("TWO = ",TWO)
	fmt.Print("\n")
	fmt.Print("BEGINS = ",BEGINS)
	fmt.Print("\n")
	fmt.Print("ONCE = ",ONE)
	fmt.Print("\n")
	fmt.Print("TWICE = ",TWICE)
}
