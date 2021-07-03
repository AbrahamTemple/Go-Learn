package main

/**
匿名与别名导包
 */

import (
	_ "../chapter5/lib1" // 匿名：虽然只调用不使用，但init依旧会被执行
	pkg2 "../chapter5/lib2" // 别名：把包里所有内容赋给pkg2
	. "../chapter5/lib2" // 别名：把包里所有内容导入当前包，直接引用即可
)

func main() {
	pkg2.TestLib()
	TestLib()
}
