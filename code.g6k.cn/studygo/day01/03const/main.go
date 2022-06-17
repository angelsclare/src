package main

import "fmt"

//常量
//常量定义后不能修改
//在程序运行期间不会改变的量

const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	netFound = 404
)

//批量声明常量如果某一行没有赋值，默认就等于上一行
const (
	n1 = 100
	n2
	n3
)

//iota 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

const (
	a1 = iota // 0
	a2 = iota // 1
	a3 = iota // 2
)

const (
	b1 = iota // 0
	b2 = iota // 1
	_         // 2
	b3 = iota // 3
)

const (
	c1 = iota
	c2 = 200
	c3 = iota
	c4
)

const (
	d1, d2 = iota + 1, iota + 2 //d1=1 d2=2
	d3, d4 = iota + 1, iota + 2 //d3=2 d4=3

)

//定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
//同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //1 2
	c, d                      //2 3
	e, f                      //3 4
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
	fmt.Println(a, b, c, d, e, f)
}
