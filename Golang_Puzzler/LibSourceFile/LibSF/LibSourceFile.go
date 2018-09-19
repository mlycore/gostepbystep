/*
package main

import "flag"

var name string

func init(){
	flag.StringVar(&name,"name","everyone","The greeting project.")
}

func main(){
	flag.Parse()
	hello(name)
}*/

package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)
}