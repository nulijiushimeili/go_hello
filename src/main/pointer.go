package main

import "fmt"

// 指针交换基本类型
func swap(a,b *int){
	*a,*b = *b,*a
}

func swap2(x,y int)(int,int){
	return y,x
}

func main(){
	var a,b = 3,4
	swap(&a, &b)
	fmt.Println(a,b)

	var c,d = 5,6
	var x,y = swap2(c,d)
	fmt.Println(x,y)
}