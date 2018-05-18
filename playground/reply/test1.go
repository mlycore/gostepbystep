package main

import (
	"fmt"
)

type Programmer interface {
	Coding(string) string
}
type Stduent struct{}

func (stu *Stduent) Coding(codes string) (comment string) {
	if codes == "happy hacking" {
		comment = "You are a talented programmer"
	} else {
		comment = "Are you ok?"
	}
	return
}
func main() {
	// var you Programer = Studuent{} 错误，
	var you Programmer = new(Stduent)
	codes := "happy hacking"
	fmt.Println(you.Coding(codes))
}
