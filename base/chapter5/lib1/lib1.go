package lib1

import "fmt"

//首字母大写才会暴露
func TestLib()  {
	fmt.Println("lib1包的方法被调用")
}

func init() {
	fmt.Println("lib1包的init被调用")
}
