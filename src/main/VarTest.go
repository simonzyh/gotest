package main

import (
	"fmt"
	"unsafe"
)

var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

func main1() {
	println(a, b, c)
	d := a
	println(&d)
	println(&b)

	println(&a)

}

func main() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = 10  //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	j := 123

	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	a1 := "hello12121212"

	fmt.Println(unsafe.Sizeof(a1), len(a1))
}
