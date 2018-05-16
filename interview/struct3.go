package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	fmt.Printf("%p\n", stu)
	return stu
}

func main() {
/*
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
*/
	p := live()	
	if p == nil {
		fmt.Println("AAAAAAA")
	}else {
		fmt.Println("BBBBBBB")
	}
	fmt.Printf("%p, %v\n", p, p == nil)
}

// wrong
// AAAAAAA

// right
// BBBBBBB
