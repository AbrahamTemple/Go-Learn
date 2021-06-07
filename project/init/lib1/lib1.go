package lib1

import "fmt"

//首字母大写才会暴露
func TestLib()  {
	fmt.Print("lib1包的方法被调用")
	fmt.Print("\n")
}

func init() {
	fmt.Print("lib1包的init被调用")
	fmt.Print("\n")
}
