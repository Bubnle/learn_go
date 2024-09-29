// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
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

	selectUse()

	sumTenLoop()

	gotoUse()

	a := 5
	b := 3
	ref(&a, &b)
	fmt.Printf("a=%d , b =%d\n", a, b)
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

	var p = 9999
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

func selectUse() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 使用go 关键字启动一个协程，执行后面的匿名函数
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		// 将c1收到的消息赋值给msg
		case msg := <-c1:
			fmt.Println("received", msg)

		case msg := <-c2:

			fmt.Println("received", msg)
			break

		}
	}
}

func sumTenLoop() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 100

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Println("key:", key, "value:", value)
	}

	for key := range map1 {
		fmt.Println("key:", key)
	}

	for _, va := range map1 {
		fmt.Println("value:", va)
	}

}

func gotoUse() {

	sum := 10
label:
	for i := 1; i <= 10; i++ {
		fmt.Println("i=", i)
		sum += i
		if sum == 16 {
			goto label
		}
	}

	fmt.Println(sum)
}

func ref(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
