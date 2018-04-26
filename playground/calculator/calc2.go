package main

import (

	"fmt"

)

func Add(a, b float32) float32 {
	return a + b
}

func Minus(a, b float32) float32{
	return a - b
}

func Multiply(a, b float32 ) float32{
	return a * b
}

func Div(a, b float32) float32{
	if b == 0{
		return 0
	}else {
		return a / b
	}
}

var(
	a float32
	b float32
	operating string
)
func main(){
	println("请输入a、操作符和b的值")
	fmt.Scanln(&a, &operating, &b)
	switch operating{
	case "+": fmt.Printf("算式的结果是 %.2f %s %.2f = %.2f\n", a,operating,b,Add(a,b))
	break
	case "-": fmt.Printf("算式的结果是 %.2f %s %.2f = %.2f\n", a,operating,b,Minus(a,b))
	break
	case "*": fmt.Printf("算式的结果是 %.2f %s %.2f = %.2f\n", a,operating,b,Multiply(a,b))
	break
	case "/": fmt.Printf("算式的结果是 %.2f %s %.2f = %.2f\n", a,operating,b,Div(a,b))
	break
	}

}