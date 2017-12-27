package main

import (
	"fmt"
)

type Phone interface {
	call()
	call1()
}
type Phone1 interface {
	call()
	call1()
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone NokiaPhone) call1() {
	fmt.Println("I am Nokia, I can call1 you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) call1() {
	fmt.Println("I am iPhone, I can call1 you!")
}
func main() {
	var phone Phone1

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}