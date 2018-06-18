// 当前程序的包名
package main

// 导入其他包
import (
	"fmt"
	"math/cmplx"
	"math"
)

/*
 *  关键字
	下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	goto	package	switch
	const	fallthrough	if	range	type
	continue	for	import	return	var


	除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	uintptr
	程序一般由关键字、常量、变量、运算符、类型和函数组成。

	程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

	程序中可能会使用到这些标点符号：.、,、;、: 和 …。
 */


// 定义常量
const PI = 3.14

// 包内全局变量的声明和赋值
var name = "go"
var aa = 3
var bb = true
// 包内全局变量的声明和赋值
var(
	cc = 5
	dd = true
)

// 一般类型声明
type newType int

// 结构的声明
type st struct{}

// 接口的声明
type golang interface {

}

// 定义变量不赋值
func variable(){
	var a int
	var b string
	fmt.Printf("%d %q",a,b)
	fmt.Println()
}

// 同时定义多个变量赋值
func variable1(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

// 可以定义不同的类型
func variable2(){
	var a,b,c,d = 1,2,true,"def"
	fmt.Println(a,b,c,d)
}

// 推荐的定义变量的形势,使用 :,在函数的外面是不能使用:的
func variable3(){
	a,b,c,d := 4,5,false,"str"
	fmt.Println(a,b,c,d)
}

// 内建变量类型
/*
    bool,string
	int,int8,int16,int32,int64,uintptr
	byte,rune   // byte是8字节, rune是32字节
	float32,float64,complex64,complex128  // complex是复数类型
 */

// 复数的测试--欧拉公式
func euler(){
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	// 欧拉公式: e^(i*pi) + 1 = 0
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi) + 1)
	// e是一个特殊的常数,欧拉公式还可以这样表示
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

// 测试三角函数
func triangle(){
	var a,b int = 3,4
	var c int
	// go语言是需要强制类型转换的(注意考虑精度问题)
	c = int(math.Sqrt(float64((a*a + b*b))))
	fmt.Println(c)
}

func consts(){
	// 一般go语言常量不大写
	const filename = "abc.txt"
	const a,b = 3,5
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)
}

// 枚举类型
func enums(){
	const(
		cpp = 0
		java = 1
		js = 2
		c = 3
		python = 4
	)
	const(
		// iota 自增值  0 1 2 3
		a = iota
		b
		e
		d
	)
	fmt.Println(cpp,c,java,js,python)
	fmt.Println(a,b,e,d)
}

// 由main函数作为程序入口点启动
func main() {
	fmt.Println("hello word!")
	variable()
	variable1()
	variable2()
	variable3()
	euler()
	triangle()
	consts()
	enums()
}
