package main

import "fmt"

func SetValueByArray(arr [5]int) {

	arr[0] = 100
	fmt.Printf("mid %p\n", &arr)
}



func SetValueBySlice(slice []int) {

	slice[0] = 100

}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("before %p\n", &arr)
	SetValueByArray(arr)
	fmt.Printf("after %p\n", &arr)

	fmt.Println(arr)

	slice := arr[:]

	SetValueBySlice(slice)

	fmt.Println(arr)

}