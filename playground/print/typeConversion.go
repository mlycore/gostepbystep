package main

import (
	"fmt"
	"strconv"
)

func main(){
	//Int转字符串
	fmt.Println(strconv.Itoa(100))

	//64位整型转字符串
	var i int64
	i = 0x100
	fmt.Println(strconv.FormatInt(i, 16))

	//字符串转Int
	fmt.Println(strconv.Atoi("97"))

	//字节转字符串
	fmt.Println(string([]byte{33, 99, 100, 66}))

	//字符串转字节
	fmt.Println([]byte("abcde"))

	//字符串转float
	s := "0.1234567890123456789"
	f, _ := strconv.ParseFloat(s, 32)
	fmt.Println(f)
	f, _ = strconv.ParseFloat(s, 64)
	fmt.Println(f)

	//float转字符串
	string := strconv.FormatFloat(f, 'E', -1,32)
	fmt.Println(string)
	string = strconv.FormatFloat(f, 255,-3,64)
	fmt.Println(string)

}

