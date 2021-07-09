package main

import "fmt"
/**
map声明应用
 */
func main() {

	//------- 声明 -----------
	var m map[string] string

	b := map[bool] string{m==nil:"指针为空",m!=nil:"指针不为空"}[true]
	fmt.Println(b)

	//分配10容量
	m = make(map[string] string, 10)

	b = map[bool] string{m==nil:"指针为空",m!=nil:"指针不为空"}[true]
	fmt.Println(b)

	m["A"] = "java"
	m["B"] = "go"
	m["C"] = "c++"

	fmt.Println(m)

	//分配可变容量
	m1 := make(map[int] string)

	m1[0] = "java"
	m1[1] = "go"
	m1[2] = "c++"

	fmt.Println(m1)

	//声明并赋值
	m2 := map[string] string{
		"one": "java",
		"two": "go",
	}

	fmt.Println(m2)


	//------- 应用 -----------
	delete(m2,"one") //删除map指定key

	for k, v := range m2 {
		fmt.Println("k=",k,"v=",v)
	}
}
