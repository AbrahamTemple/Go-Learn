package main

/**
init与导包
 */

// 请将GO111MODULE设为自动: go env -w GO111MODULE=auto

import (
	"../chapter5/lib1"
	"../chapter5/lib2"
)

func main()  {
	// 所有人的init会最先被执行完
	lib1.TestLib()
	lib2.TestLib()
}
