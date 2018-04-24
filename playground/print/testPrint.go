package main

import (
	"fmt"
	"strconv"

)

func main() {
	var a int = 9
	var b float32 = 5
	fmt.Println(strconv.Itoa(a))
	//var c = float32(a)
	fmt.Println("a is %f",a)
	fmt.Printf("b is %.3f",b)
	//fmt.Printf("c is %.6f，b is %.3f\n",c,b )

	//练习Int转字符串
	fmt.Println(strconv.Itoa(100))

	//练习字符串转Int
	//fmt.Println(strconv.Atoi("97")),之前感觉为啥不像Itoa那样呢，好吧写了后知道了hah
	i, _ := strconv.Atoi("97")
	fmt.Println(i)

	//练习字节转字符串
	fmt.Println(string([]byte{98, 97, 99, 100}))

	//练习字符串转字节
	fmt.Println([]byte("abcd"))


	//练习64位整形转字符串
	var j int64
	j = 0x100
	fmt.Println(strconv.FormatInt(j, 10))

	//练习字符串转到float
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568


	//练习float转到字符串

	string := strconv.FormatFloat(f,'E', -1, 32)
	fmt.Println(string)
	string = strconv.FormatFloat(f,'E', -1, 64)
	fmt.Println(string)

	/*//字节转32位整型
	d := []byte{0x00, 0x00, 0x03, 0xe8}
	//fmt.Println(d) //就是想看看输出的是啥
	bytesBuffer := bytes.NewBuffer(d)
	//fmt.Println(bytesBuffer)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	fmt.Println(x)

	//32位整型转字节
	var y int32
	y = 106
	bytesBuffer1 := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer1, binary.BigEndian, y)
	fmt.Println(bytesBuffer1.Bytes())*/
}

