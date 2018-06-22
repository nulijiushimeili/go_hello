package main

import "fmt"

/**
统计一个单词中的不重复的字符的个数
 */

 /*
 有问题待修复
  */
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}



func main() {
	fmt.Println("hello:", lengthOfNonRepeatingSubStr("hello"))
	//fmt.Println("golang:", lengthOfNonRepeatingSubStr("golang"))
	//fmt.Println("python:", lengthOfNonRepeatingSubStr("python"))
	//fmt.Println("hadoop:", lengthOfNonRepeatingSubStr("hadoop"))
	//fmt.Println("java:", lengthOfNonRepeatingSubStr("java"))
}
