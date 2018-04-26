package main
//读取文本文件

import (
	"io/ioutil"
	"fmt"
)

func main() {
	b, err := ioutil.ReadFile("/Users/yp-tc-m-2642/Desktop/hello")
	if err != nil {
		fmt.Print(err)
	}
	/*fmt.Println(b)//返回的是字节型数组
	str := string(b)//将字节型数组转换成字符串
	fmt.Println(str)
	*/
	//上面的标注内容可以通过下面的方式实现
	fmt.Printf("%s\n", b)
}
