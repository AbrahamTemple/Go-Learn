package lib2

import "fmt"

//首字母大写才会暴露
func TestLib()  {
	fmt.Print("lib2包的方法被调用")
	fmt.Print("\n")
}

func init() {
	fmt.Print("lib2包的init被调用")
	fmt.Print("\n")
}
