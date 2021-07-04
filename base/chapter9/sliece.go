package main

import "fmt"

/**
数组切片
 */

func main() {
	var arr []int
	fmt.Println(arr,"数组是否为空:",arr==nil) //数组为空，存储任何下标位都会报错
	arr = make([]int,3) //初始化出3个位置值为0的int数组
	fmt.Println(arr,"数组是否为空:",arr==nil)

	nums := make([]int,2,4) //2个位置，4个容量(容量不是长度，而是追加限度)
	fmt.Println("长度:", len(nums),"容量:",cap(nums),"数组:",nums)

	nums = append(nums, 1)//追加1个1
	fmt.Println("长度:", len(nums),"容量:",cap(nums),"数组:",nums)

	nums = append(nums, 2)//追加一个2，此时容量已达到定义的极限
	fmt.Println("长度:", len(nums),"容量:",cap(nums),"数组:",nums)

	nums = append(nums, 3)//追加一个3，开辟新的4个容量跟原来的4容器合并
	fmt.Println("长度:", len(nums),"容量:",cap(nums),"数组:",nums)


	fmt.Println("nums切片:",nums[2:5])

	ns := nums //ns和nums都指向同一个地址
	ns[3] = 0

	fmt.Println("ns切片:",ns[2:5])
	fmt.Println("nums切片:",nums[2:5])

	ns2 := make([]int,len(nums),cap(nums))
	copy(ns2,nums) //将nums拷贝到ns，ns用新的地址

	ns2[3] = 8
	fmt.Println("ns2切片:",ns2[2:5])
	fmt.Println("nums切片:",nums[2:5])
}
