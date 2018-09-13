package main

import (
	"flag"
	"fmt"
)

func main(){
	/*var name string //[1]
	flag.StringVar(&name,"name","everyone","The greeting project.") //[2]*/

	/*//[1]&[2] to one code
	var name = *flag.String("name","everyone","The greeting project.")*/

	// use the short var declaration
	name := *flag.String("name","everyone","The greeting project.")

	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
