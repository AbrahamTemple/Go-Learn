package lib2

import "fmt"

//首字母大写才会暴露
func TestLib()  {
	fmt.Println("lib2包的方法被调用")
}

func init() {
	fmt.Println("lib2包的init被调用")
}
