package main
import (
	"fmt"
)
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		// m[stu.Name] = &stu
		s := stu
		m[s.Name] = &s
		// fmt.Printf("%p, %s\n", &stu, stu.Name)
		fmt.Printf("%p, %s\n", &s, s.Name)
	}
	fmt.Printf("%v\n", m)	
}

func main() {
	pase_student()
}
