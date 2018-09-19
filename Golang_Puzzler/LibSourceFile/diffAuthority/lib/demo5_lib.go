package lib

import (
	"os"
	in "gostepbystep/Golang_Puzzler/LibSourceFile/diffAuthority/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}