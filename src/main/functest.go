package main
import "fmt"

/* 函数返回两个数的最大值 */
func max(num1, num2 *int) int {
	/* 声明局部变量 */
	var result int

	if (*num1 > *num2) {
		result = *num1
		*num1 = *num2;
		*num2 = result;
	} else {
		result = *num2
		*num2 = *num1;
		*num1 = result;
	}

	return result
}

func main12() {
	/* 定义局部变量 */
	var a int = 300
	var b int = 400
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(&a, &b)

	fmt.Printf("最大值是 : %d\n", ret)
	fmt.Printf("", a)

}