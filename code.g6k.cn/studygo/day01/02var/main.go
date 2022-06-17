package main

import "fmt"

//申明变量

// var name string
// var age int
// var isok bool

//批量申明（注意非全局变量如果申明未使用编译时就会报错）
var (
	name string
	age  int
	isok bool
)

func main() {
	name = "yx"
	age = 16
	isok = true
	fmt.Print(isok)               //打印结束不会换行
	fmt.Printf("name:%s\n", name) //$s占位符
	fmt.Println(age)              //打印并换行

	//申明变量同时并赋值
	var s1 string = "大师"
	fmt.Printf("帅逼：%s", s1)

	//类型推导（根据值判断变量类型是什么类型）
	var s2 = "20"
	fmt.Println(s2)

	//简短变量声明
	s3 := "大师帅逼"
	fmt.Println(s3)
}
