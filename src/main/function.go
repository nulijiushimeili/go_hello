package main

import "fmt"

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
}
