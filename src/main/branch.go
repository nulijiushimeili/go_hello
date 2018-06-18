package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// if 分支的使用
func testif(){
	const filename = "file\\abc.txt"
	contents,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("%s\n",string(contents))
	}
}


// if 还可以这样写
func testif2(){
	const filename = "file\\abc.txt"
	if contents,err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else{
		// 返回的是字节数组,
		fmt.Println(contents)
	}
}

// switch 分支的使用
func grade(score int)string{
	g :=  ""
	switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 100:
		g = "A"
	default:
		// 终端程序的执行
		panic(fmt.Sprintf("Wrong score: %d",score))
	}
	return g
}

// for 循环 (将传进来的整数,转换成二进制)
func convertToBin(n int)string{
	res := ""
	for ; n>0 ;n/=2{
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

// read file by line
func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

/*
	for,if 后面的条件没有()
	if 条件里也可以定义变量
	没有while
	switch 不需要break,也可以直接switch多个条件
 */
func main(){
	testif()
	testif2()

	//fmt.Println(
	//	grade(101),
	//	grade(81),
	//	grade(28),
	//	grade(56),
	//	grade(90),
	//)

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("file\\abc.txt")

}
