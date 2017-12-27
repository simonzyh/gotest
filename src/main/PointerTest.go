package main

import "fmt"

func main() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Printf("ptr 的值为 : %x\n", ptr == nil)

	a := []int{10, 100, 200}
	var i int
	var ptr1 [3]*int;

	for i = 0; i < 3; i++ {
		ptr1[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr1[i])
	}
	fmt.Printf("a", a)

}