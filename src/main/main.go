package main

import (
	"cal"
	"fmt"
	"cal/multi"
)

func main() {
	result := cal.Add(1, 2)
	fmt.Println(result)
	result = multi.Multi(2, 5)
	fmt.Println(result)

	var a int = 300
	var b int = 400
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(&a, &b)
	fmt.Println(ret)
}