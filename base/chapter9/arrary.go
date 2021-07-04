package main

import (
	"fmt"
)

/**
数组声明与循环
 */

func main() {
	var arr1 [2]int //默认值为0

	arr2 := [5]int{1,2,3} //不填充部分默认值为0

	arr3 := []float32{0.1,0.2,0.3} //动态长度

	x,y := 1, 2

	arr4 := [...]*int{&x,&y} //指针数组

	var arr5 = [...]int{1:'a',2:'b'} //字符自动转义成数字

	fmt.Println("第一类循环，arr1") //自动匹配长度
	for i := range arr1 {
		fmt.Println("value =" ,arr1[i])
	}

	fmt.Println("第二类循环，arr2") //附带index
	for index, value := range arr2 {
		fmt.Println("index =" ,index,"value =" ,value)
	}

	fmt.Println("第三类循环，arr3") //index匿名
	for _, value := range arr3 {
		fmt.Println("value =" ,value)
	}


	change(&x,&y)
	//doForArray(&arr4)
	//fmt.Println("默认循环，arr4")
	//for i := 0; i< len(arr4); i++ {
	//	fmt.Println("index =" ,i,"value =" ,arr4[i])
	//}

	fmt.Println("默认循环，arr4")
	for i := 0; i< len(arr4); i++ {
		fmt.Println("index =" ,i,"address =" ,arr4[i],"value =",*arr4[i])
	}

	fmt.Println("第二类循环，arr5") //0号位没初始化，则值为0
	for index, value := range arr5 {
		fmt.Println("index =" ,index,"value =" ,value)
	}
}

func change(px *int,py *int)  {
	*px = 25
	*py = 36
}

func doForArray(arr *[]int) {

	//指针数组
	for index, value := range *arr {
		fmt.Println("index =" ,index,"value =" ,value)
	}

}
