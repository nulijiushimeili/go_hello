package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// +-*/ 的计算
func eval(a,b int, op string) int{
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a - b
	default:
		panic("unsupported operation:" + op)
	}
}

func eval2(a,b int, op string)(int,error){
	switch op{
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0,fmt.Errorf(
			"unsupport operation:%s",op,
		)
	}
}

func div(a,b int) (x,y int){
	return a*b, a/b
}

// 不推荐
func div2(a,b int)(x,y int){
	x = a * b
	y = a / b
	return
}

// go也是函数式编程语言,它的函数式一等公民
// 函数可以作为参数传递,函数里面也可以再写别的函数
func apply(op func(int,int) int, a, b int)int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf(
		"Calling function %s with args " +
			"(%d, %d): ",opName,a,b,
		)
	return op(a,b)
}

func pow(a,b int)int{
	return int(math.Pow(float64(a),float64(b)))
}

func sum(nums... int)int{
	s := 0
	for i := range nums{
		s += nums[i]
	}
	return s
}

/*
函数多个返回值时,可以起名字
对于调用者没有区别
 */
func main(){
	fmt.Println(eval(3,4,"*"))
	a, b := div(30,6)
	fmt.Println(a,b)


	if result,err := eval2(3,4,"x");err != nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println(result)
	}


	res1 := apply(pow,4,5)
	fmt.Println(res1)

	// 匿名函数
	fmt.Println(apply(
		func(a,b int)int{
			return int(math.Pow(float64(a),float64(b)))
		},3,4,
	))

	// 无限制参数
	fmt.Println(sum(1,2,3,4,5))
}
