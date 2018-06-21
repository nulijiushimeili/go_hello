package main

import "fmt"

/*
slice, 对数组的切片
slice 是对 array 的 view
slice 是可以向后扩展的,只要不超过slice的capacity
slice.capacity()是从切片的开始到数组的结束的的长度
 */
func printArray(arr *[5]int){
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

func updateSlice(s []int){
	s[0] = 100
}

func main(){
	arr := [...]int {0,1,2,3,4,5,6,7}

	fmt.Println("arr[2:6] = ",arr[2:6])
	fmt.Println("arr[:6] = ",arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ",arr[2:])
	s2 := arr[:]
	fmt.Println("arr[:] = ",arr[:])

	fmt.Println("After updateSlice s1...")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice s2 ...")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice ...")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[:2]
	fmt.Println(s2)

	fmt.Println("slice 是可以向后扩展的,只要不超过slice的capacity")
	s2 = arr[2:6]
	fmt.Println(s2)
	s3 := s2[3:5]
	fmt.Println(s3)


}