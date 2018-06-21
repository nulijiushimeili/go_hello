package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func newSlice() {
	fmt.Println("new Slice ...")
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copy slice ...")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Delete elements from slice ...")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front ... ")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back ...")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	// s4,s5,已经不是arr的视图了,是系统重新copy的新的更大的数组
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("slice append()")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(arr)

	var s []int
	//s = nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	newSlice()

}
