package main
import (
	"fmt"
)
func main() {
	s := make([]int, 5)
	fmt.Printf("before append: %d, %d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Printf("after append: %d, %d\n", len(s), cap(s))

	slice := []int{1, 2, 3, 4}
	fmt.Printf("slice before append: %d, %d\n", len(slice), cap(slice))
	newSlice := append(slice, 5)
	fmt.Printf("newslice after append: %d, %d\n", len(newSlice), cap(newSlice))

}

//t: 1 2 3 0 0
//m: 0 0 0 0 0 1 2 3 0 0
//right: 0 0 0 0 0 1 2 3
