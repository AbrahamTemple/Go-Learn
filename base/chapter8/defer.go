package main

import "fmt"

/**
反序执行
 */

// 再怎么样，return关键字都比defer关键字更后执行

func process1()  {
	fmt.Println("---1号程序被执行---")
}

func process2()  {
	fmt.Println("---2号程序被执行---")
}

func process3() string {
	defer process1()
	defer process2()
	return "---3号程序被执行---"
}

func main() {
	//defer: 程序是从上往下先后执行变成从下往上先后执行
	defer process1()
	defer process2()
	fmt.Println(process3())
}
