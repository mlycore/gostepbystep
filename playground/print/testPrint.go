package main

import "fmt"

func main() {
	var a int = 4
	var b float32 = 5
	var c = float32(a)
	fmt.Println("a is %f",a)
	fmt.Printf("c is %.6f，b is %.3f\n",c,b )
}
