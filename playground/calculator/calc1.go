package main

import (

	"fmt"
)

var (
	num1 float32
	num2 float32
	operation string
	f float32

)
func main() {
	fmt.Println("Please input your number（number operation number）: ")
	fmt.Scanln(&num1,&operation,&num2)
	if operation == "+"{
		f = num1 + num2
	} else if operation == "*"{
		f = num1 * num2
	}else if operation == "-" {
		f = num1 - num2
	}else if operation == "/" {
		f = num1 / num2
	}
	fmt.Printf("%f %s %f = %f\n", num1, operation, num2, f)

}


//浮点型转字符串、 小程序：计算器 2位的 ，输入
//特殊情况处理 不能除以0,这个没实现