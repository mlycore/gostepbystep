package main

import(
	"flag"
	"gostepbystep/Golang_Puzzler/LibSourceFile/diffCodePackage/demo4_lib"
)


var name string

func init(){
	flag.StringVar(&name,"name","everyone","The greeting product.")
}

func main(){
	flag.Parse()
	demo4_lib.Hello(name)
}
