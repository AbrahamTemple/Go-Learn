package main

import "fmt"
/**
指针
 */

// 指针就是地址，地址就是指针
// &代表地址，*代表地址数据结构，**代表地址的地址的数据结构（上一层必须是地址）

func main() {
	var a int = 1
	change1(a) //这时的指针指针指向形参的地址，没能改变a原内存里的值
	fmt.Println(a)
	change2(&a) //让指针指向原内存地址，并且改变内存里的值
	fmt.Println(a)

	var b  = 2
	switches(&a,&b) //互换a和b的值
	fmt.Println("a =",a,"b =",b)

	var pa *int
	pa = &a //a的地址赋给pa数据结构
	fmt.Println(pa)

	var ppa  ** int //二层地址
	ppa = &pa //pa数据结构的地址赋给ppa数据结构

	fmt.Println(ppa)
}

func change1(p int)  {
	p = 10
}

func change2(p *int)  {
	*p = 10
}

func switches(a *int, b *int)  {
	var t int
	t = *b //b原内存值给t
	*b = *a //a原内存值给b
	*a = t //b值给a作为原内存值
}
