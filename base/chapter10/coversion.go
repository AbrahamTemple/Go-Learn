package main

import "strconv"
import "fmt"

/**
类型强制转换
 */

const (
	str = "6"
	num = 8
)

func main() {

	//string转成int：
	strnum, err1 := strconv.Atoi(str)
	fmt.Println("string转成int=",strnum,"异常",err1)
	//string转成int64：
	strnum64, err2 := strconv.ParseInt(str, 10, 64)
	fmt.Println("string转成int64=",strnum64,"异常",err2)
	//int转成string：
	numstr := strconv.Itoa(num)
	fmt.Println("int转成string=",numstr)

	//int转化为int64
	int64num := int64(num)

	//int64转成string：
	num64str := strconv.FormatInt(int64num,10)
	fmt.Println("int64转成string=",num64str)
}
