package main

import "fmt"

func main(){
	arr := [...]int{1, 2, 3, 4, 5}

	len := len(arr)   //len获取数组长度

	fmt.Println("修改前:", arr)

	arr[0] = 100      //下标访问数组元素

	fmt.Println("修改后:", arr)

	fmt.Println("数组长度:", len)
}
