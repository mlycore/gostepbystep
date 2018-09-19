package main

import (
	"flag"
	//in "gostepbystep/Golang_Puzzler/LibSourceFile/diffAuthority/lib/internal" // 此行无法通过编译。
	//"os"
	"gostepbystep/Golang_Puzzler/LibSourceFile/diffAuthority/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}