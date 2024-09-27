// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// main 函数
func main() {
	var s, sep string
	// fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	Out()

	formatUse()

	variety()

	pointer()

	iotaUse()

	condition(-10)

	switchUse(1)
}

/*
	格式化符号

这里面的Sprintf 就是相当于 c++ 的printf
%p 指针以十六进制显示
%v 按值的本来值输出
%+v 在%v的基础上，对结构体字段名和值进行展开
%#v 输出Go语言的语法格式的值
%T 输出Go语言格式的类型和值
%% 输出%
*/
type point struct {
	x, y int
}

func formatUse() {
	var x string
	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	x = fmt.Sprintf("%T\n", p)
	fmt.Println(x)
	var stockCode = 123
	var endDate = "2024-9-27"
	var url = "Code=%d&endDate=%s"
	var target = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(target)
}

func variety() {
	var pp, yy int
	pp = 12
	fmt.Println(yy)
	fmt.Println(pp)

	var p int = 9999
	fmt.Println(p)
}

func pointer() {
	var x = 10
	a := &x
	fmt.Println(a)
	x = 100
	fmt.Println(a)

	y := 1000
	var po *int
	po = &y
	fmt.Println(*po)
}

func iotaUse() {
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i, j, k, l)
}

func condition(a int) {
	if a < 10 {
		fmt.Println(a)
	} else {
		fmt.Println("wrong")
	}
}

func switchUse(a int) {
	switch a {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2", "force execute")

	default:
		fmt.Println("is not 1 or 2")
	}
}
