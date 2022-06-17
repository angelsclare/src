package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"E:\\go\\src\\code.g6k.cn\\studygo\\day01\\09str\""
	fmt.Printf("代码路径是：%s", path)
	//定义多行的字符串
	s2 := `
		是轻薄
		人情恶
		与送黄昏花易洛
	`
	fmt.Println(s2)
	s3 := `E:\go\src\code.g6k.cn\studygo\day01\09str`
	fmt.Println("反引号就不需要转义：", s3)

	/*
	   len(str)								求长度
	   +或fmt.Sprintf						拼接字符串
	   strings.Split						分割
	   strings.contains						判断是否包含
	   strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	   strings.Index(),strings.LastIndex()	子串出现的位置
	   strings.Join(a[]string, sep string)	join操作
	*/

	fmt.Println(len(s3))

	name := "大师"
	world := "帅逼"

	fmt.Println(name + world)
	ss1 := fmt.Sprintf("%s%s", name, world) //Sprintf返回字符串变量
	//fmt.Println("%s%s", name, world)   //Println可以打印出字符串，和变量

	fmt.Printf(ss1)
	fmt.Println(ss1)

	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s3, "你好"))
	fmt.Println(strings.Contains(s3, "ode.g6k.cn"))
	//前缀和后缀
	fmt.Println(strings.HasPrefix(s3, "E:"))
	fmt.Println(strings.HasPrefix(s3, "09str"))
	fmt.Println(strings.HasSuffix(s3, "09str"))
	fmt.Println(strings.HasSuffix(s3, "E:"))
	//位置
	fmt.Println(strings.Index(s3, "o"))
	fmt.Println(strings.LastIndex(s3, "o"))

	//拼接
	fmt.Println(strings.Join(ret, "+")) //分割后拼接

}
