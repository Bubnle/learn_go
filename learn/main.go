// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Circle struct {
	radius float32
}

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

	showSequence()

	anonymous()

	c := Circle{1}
	showArea(c)

	scopeOfEffectUse()

	arrayUse()

	pointerUse()

	SliceUse()

	rangeUse()

	mapUse()

	fmt.Println(fibonacci(10))

	transform()

	interfaceUse()

	interfaceUse2()

	transInterface()

	transInterfaceUse2()

	ErrorUse()

	goroutineUse()

	ChannelUse()

	ChannelUse2()

	ChannelUse3()

	SelectUse()
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

	// go 关键字使用 跳到函数 //
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

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func showSequence() {
	nextN := getSequence()
	fmt.Println(nextN())
	fmt.Println(nextN())
	fmt.Println(nextN())

	nextN1 := getSequence()
	fmt.Println(nextN1())

}

func anonymous() {
	add := func(a, b int) int {
		return a + b
	}
	result := add(1, 2)
	fmt.Println("1+2=", result)

	multiply := func(a, b int) int {
		return a * b
	}
	product := multiply(1, 2)
	fmt.Println("1*2=", product)

	calculate := func(operation func(int, int) int, a, b int) int {
		return operation(a, b)
	}

	sum := calculate(add, 2, 8)

	fmt.Println("匿名函数的结果是", sum)

	mul := calculate(multiply, 2, 8)

	fmt.Println("匿名函数乘法的结果", mul)

}

func (c Circle) getArea() float32 {
	return 3.14 * c.radius * c.radius
}

func showArea(c Circle) {
	fmt.Println(c.getArea())
}

func scopeOfEffectUse() {
	var a, b int // 这个是局部变量
	a = 2
	b = 2
	add := func(a, b int) int { // 这里面的a b是函数内定义
		return a + b
	}

	fmt.Println(add(a, b))
}

func arrayUse() {
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array[1])
	fmt.Println("array=", array)
	balance := [10]string{"1", "hello"}
	fmt.Println(balance[1])

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	bin := [...]int{99, 88, 33, 22, 21, 111}
	i := 1
	for i < len(bin) {
		fmt.Println(bin[i])
		i++
	}
	// 二维
	var n = [5][5]int{{1}, {2}, {3}, {4}}

	fmt.Println(n[0][0])

	modify := func(arr []int, size int) {
		for i := 0; i < len(arr); i++ {
			arr[i] *= 2
		}
	}

	temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	modify(temp, 10)

	fmt.Println("值传递之后的", temp)

	modify2 := func(arr *[]int, size int) {
		for i := 0; i < len(*arr); i++ {
			(*arr)[i] *= 2
		}
	}
	modify2(&temp, 10)

	fmt.Println(`传指针`, temp)
}

func pointerUse() {
	x := 10
	p1 := &x
	fmt.Println(*p1)

	var p2 *int
	p2 = &x
	fmt.Println(*p2)

	a := []int{1, 2, 3, 4, 5}
	var ptr [3]*[]int // 每个指针是指向切片的
	for i := 0; i < len(ptr); i++ {
		ptr[i] = &a
	}
	fmt.Println("指针数组是", ptr)

	fmt.Println((*ptr[0])[0], *ptr[1], *ptr[2])

	fmt.Println("以下是指向指针的指针")
	var pp int
	var pp1 *int
	var pp2 **int

	pp = 3000

	pp1 = &pp
	pp2 = &pp1

	fmt.Println(pp)
	fmt.Println(pp1, *pp1)
	fmt.Println(pp2, *pp2, **pp2)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func SliceUse() {
	arr := []int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(s)

	s1 := arr[:]
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)

	// 不包括2下标
	s3 := arr[0:2]
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

	n := func() {
		var number []int
		fmt.Println(number, len(number), cap(number))
		if number == nil {
			fmt.Println("this is nil")
		}
	}
	n()

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array[1:10]", array[1:10])
	fmt.Println("array[2:5]", array[2:5])
	fmt.Println("array[:10]", array[:10])

	appendAndCopyUse := func() {
		var numbers []int
		printSlice(numbers)
		// append nil slice
		numbers = append(numbers, 0)
		printSlice(numbers)

		numbers = append(numbers, 1, 2, 3)
		printSlice(numbers)

		numbers1 := make([]int, len(numbers), cap(numbers)*2)

		copy(numbers1, numbers)
		printSlice(numbers1)
	}
	appendAndCopyUse()
}

func rangeUse() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3}
	str := "Hello"

	for key, value := range arr {
		fmt.Println(key, value)
	}
	for _, value := range s {
		fmt.Println("value:", value)
	}
	for index := range str {
		fmt.Println("index:", index)
	}
	ss := "hello"
	for i, c := range ss {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}

func mapUse() {
	//var m map[string]string
	m := make(map[string]string)
	m["Google"] = "谷歌"
	m["Baidu"] = "百度"
	m["Wiki"] = "维基"

	for x := range m {
		fmt.Println(x, "中文是:", m[x])
	}

	name, name2 := m["Google"]
	fmt.Println(name, name2)

	n1, n2 := m["China"]
	fmt.Println(n1)
	fmt.Println(n2)

	cC := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New delhi",
	}
	for c := range cC {
		fmt.Println(c, "首都是", cC[c])
	}
	delete(cC, "France")
	fmt.Println("法国条目被删除")

	for c := range cC {
		fmt.Println(c, "首都是", cC[c])
	}
}

func fibonacci(a int) int {
	if a < 2 {
		return a
	}
	return fibonacci(a-1) + fibonacci(a-2)
}

func transform() {
	a := 1.222
	fmt.Println(int(a))
	x := "9999"
	fmt.Println(strconv.Atoi(x))
	n1, n2 := strconv.Atoi("name")
	fmt.Println(n1, n2)
}

type Phone interface {
	call() // 定义的接口
}
type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am NokiaPhone, i call you")
}

type iPhone struct {
}

func (iPhone iPhone) call() {
	fmt.Println("I am iPhone, i call you")
}
func interfaceUse() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(iPhone)
	phone.call()
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}
type Cir struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Cir) area() float64 {
	return c.radius * c.radius * math.Pi
}

func interfaceUse2() {
	var shape Shape
	shape = Rectangle{1, 1}
	fmt.Println(shape.area())
	shape = Cir{2}
	fmt.Println(shape.area())
}

// 这里面的i 是一个接口
func transInterface() {
	var i interface{} = "Hello"
	s, ok := i.(string)
	if ok {
		fmt.Println("success", s)
	} else {
		fmt.Println("转换失败")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}
type StringWriter struct {
	str string
}

func (sw StringWriter) Write(p []byte) (int, error) {
	sw.str = string(p)
	return len(p), nil
}
func transInterfaceUse2() {
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, World"

	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)
}

type DivideError struct {
	dividend int
	divider  int
}

// 实现的是error接口中的Error函数
func (de DivideError) Error() string {
	/// 格式化字符串
	strFormat := `
    Cannot proceed, the divider is zero.
    dividend: %d
    divider: 0
`
	// 这里面的%d 就是占位符
	return fmt.Sprintf(strFormat, de.dividend)
}
func Divide(varDividend int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		// 实例化 这个除法错误的结构体
		dData := DivideError{
			dividend: varDividend,
			divider:  varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividend / varDivider, ""
	}
}

func ErrorUse() {
	if result, errorMsg := Divide(100, 20); errorMsg == "" {
		fmt.Println("100/20 = ", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("!!! ", errorMsg)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func goroutineUse() {
	go say("world")
	go say("thread")
	say("hello")
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把sum发送到通道里面
}

func ChannelUse() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func ChannelUse2() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
	c <- 3
	// 此时就会报错
	// c <- 4
}
func fibonacciNew(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func ChannelUse3() {
	c := make(chan int, 10)
	go fibonacciNew(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciNew2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// 必须加！！！
			// 当 quit 通道接收到信号时，
			//case <-quit 会被执行，打印 "quit" 后立即 return 退出函数。
		}
	}
}

func SelectUse() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciNew2(c, quit)
}
