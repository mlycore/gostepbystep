package main

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	if 5.1 == Add(3.55, 1.55){
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
}

func TestMinus(t *testing.T) {
	if 6 == Minus(12.5,6.5){
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
}

func TestMultiply(t *testing.T) {
	if 16 == Multiply(2,8){
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
}

func TestDiv(t *testing.T) {
	if 0 == Div(12, 0){
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
	if 3 ==Div(6,2){
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}

}

