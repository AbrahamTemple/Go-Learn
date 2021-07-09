package main

import "fmt"

/**
继承
 */

type Animal struct {
	name string
	shout string
}

func (this *Animal) setShout(shout string) {
	this.shout = shout
}

func (this Animal) getShout() string {
	return this.shout
}

type Dog struct {
	Animal //代表继承
}

func (this *Dog) setShout(shout string) {
	this.shout = shout
}

type Cat struct {
	Animal //代表继承
}

func (this *Cat) setShout(shout string) {
	this.shout = shout
}

func main() {
	dog := Dog{Animal{name:"史努比",shout:"汪汪汪"}}
	fmt.Println(dog.shout)

	cat := Cat{Animal{name:"加菲猫",shout:"喵喵喵"}}
	fmt.Println(cat.shout)
}
