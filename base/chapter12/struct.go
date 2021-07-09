package main

import "fmt"

/**
结构体
 */

//方法或结构体如果是大写则代表public
//方法或结构体如果是小写则代表private
type User struct {
	id int
	name string
}

func (this User) getName() string {
	return this.name
}

func (this *User) setName(name string) {
	this.name = name
}

func main() {
	var u User
	u.id = 1
	u.name = "张三"

	fmt.Printf("%v\n",u)

	change(&u,0,"李四")

	fmt.Printf("%v\n",u)

	//---------实际用法----------
	u2 := User{id:2,name:"哈哈"}
	u2.setName("王五")
	fmt.Println(u2.getName())
}

func change(user *User,id int,name string){
	user.id = id
	user.name = name
}
