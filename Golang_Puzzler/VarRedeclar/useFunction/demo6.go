package main

import (
	"flag"
	"fmt"
)

func getTheflag() *string{
	return flag.String("name","everyone","The greeting project.")
}

/*// the function getTheFlag is also write to this.
func getTheFlag() *int{
	return flag.Int("name",2,"The greeting project")
}*/

func main(){
	var name = getTheflag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n",*name)
}
