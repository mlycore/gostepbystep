package main


//test the import.***. we can use the other package's variable. and can use the printf（）directly
import (

	. "gostepbystep/Golang_Puzzler/VarScope/importPackage"
	. "fmt"
    )

//var block = "package"

func main(){
	//block :="function"
	{
		block :="inner"
		// the code use var that always look for the inner scope.
		Printf("The inner block is %s!\n",block)
	}
	// the var for this block that always the inner of the block, this block din't contain child block。
	Printf("The block is %s.\n",Block)
}
