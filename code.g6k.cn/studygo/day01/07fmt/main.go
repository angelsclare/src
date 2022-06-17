package main

import (
	"fmt"
)

func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n) //查看变量的类型
	fmt.Printf("%v\n", n) //查看变量的值万能
	fmt.Printf("%b\n", n) //占位符%b表示二进制
	fmt.Printf("%d\n", n) //查看10进制
	fmt.Printf("%o\n", n) //八进制  以0开头
	fmt.Printf("%x\n", n) //十六进制 以0x开头
	var s = "大师帅逼"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s)
}
