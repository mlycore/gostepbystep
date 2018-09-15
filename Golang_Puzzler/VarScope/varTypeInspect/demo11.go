package main

import "fmt"

var container = [] string{"one","two","three"}

func main(){
	container := map[int]string{0:"one",1:"two",2:"three"}

	// method 1
	_,ok1 := interface{}(container).([]string)
	_,ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2){
		fmt.Printf("Error: unsuported container type: %T\n",container)
		return
	}
	fmt.Printf("The element is %q. (container type : %T)\n",container[1],container)


	//method 2
	elem, err := getElement(container)
	if err != nil{
		fmt.Printf("Error: %s.\n",err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",elem,container)

	/*
	fmt.Printf("the element is %q.\n", container[0] ) // the element is "one".
	fmt.Printf("the element is %s.\n", container[0] ) // the element is one.*/
}

func getElement(container1 interface{})(elem string,err error){
	switch t :=container1.(type) {
	case []string:
		elem = t[1]

	case map[int]string:
		elem = t[0]

	default:
		err = fmt.Errorf("unsuported container type : %T",container1)
		return
	}
	return
}

// the element is "one".
// the element is one.
