package main

import "fmt"

func main() {
	var slice1 [] int = make([] int,2,3)
	slice2 :=[] int {2,2,2}

	fmt.Printf("%p--%d\n",slice1,slice1)
	slice1=append(slice1,1)    //追加单个元素
	fmt.Printf("slice1:%p--%d\n",slice1,slice1)
	slice3 :=append(slice1,slice2...) //追加另一个切片
	fmt.Printf("slice3:%p--%d\n",slice3,slice3)
	slice4 := make([] int,len(slice3))
	copy(slice4,slice3)         //拷贝slice3
	fmt.Printf("copy：slice3:%p--%d\n",slice3,slice3)
	fmt.Printf("slice4:%p--%d\n",slice4,slice4)
}