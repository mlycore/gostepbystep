package main

import (
	"io"
	"os"
	"fmt"
)

func main(){
	var err error
	n,err := io.WriteString(os.Stdout,"Hello,everyone!")
	if err != nil {
		fmt.Printf("Error: %v\n",err)
	}
	fmt.Println()
	fmt.Printf("%v bytes were written.\n",n)
}
