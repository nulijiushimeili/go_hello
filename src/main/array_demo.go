package main

import (
	"fmt"
)

func foreach(arr [5]int){
	arr[0] = 100
	for _,v := range arr {
		fmt.Println(v)
	}
}

// 数组的引用传递
func arrPrint(arr *[3]int){
	arr[0] = 100
	fmt.Println(arr)
}

// 数组是值类型
// 在其他语言中,将数组传递进函数式引用传递,会改变数组中的值
// 在go语言中,传递的是数组的copy,不会改变原有数组的值
func main(){
	var arr1 [5]int
	arr2 := [3]int{4,5,6}
	arr3 := [...]int{2,3,4,0,6}
	var grid [4][5]bool

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	// 遍历
	for i := 0; i < len(arr3);i++{
		fmt.Println(arr3[i])
	}

	fmt.Println("遍历arr2---------------------------------")

	for i := range arr2{
		fmt.Println(arr2[i])
	}

	fmt.Println("遍历arr2--------------------------------")

	// 只要值,不需要下标的话,可以使用 _
	for _,v := range arr2{
		fmt.Println(v)
	}

	fmt.Println("go array 的值传递------------------------")
	foreach(arr3)
	fmt.Println("arr3 数组并没有改变-----------------------")
	fmt.Println(arr3)
	//foreach(arr2)
	fmt.Println("使用指针传递数组引用-----------------------")
	arrPrint(&arr2)
	fmt.Println("查看原来的值是否会改变---------------------")
	fmt.Println(arr2)


}
