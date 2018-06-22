package main

import "fmt"

/*
创建:make(mam[string]int)
获取元素:m[key]
key不存在时,获得value类型的初始值
用value,ok := m[key]来判断是否存在key
delete(),删除集合中的元素
使用range遍历key,或者遍历key-value
不保证遍历的顺序,如果需要顺序输出,需要手动进行排序
 */

/*
map中可以使用那些类型
map使用hash表,必须可以比较相等
除了slice,map,func等内建类型,都可以作为key
struct类型不包含上述字段,也可以作为key
 */
func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 = empty map

	var m3 map[string]int // m3 = nil

	fmt.Println(m, m2, m3)

	fmt.Println("traversing map ....")
	for k, v := range m {
		fmt.Println(k, v)
	}

	cause := m["cause"]
	fmt.Println("cause: ", cause)

	if qualy, ok := m["cause"]; ok {
		fmt.Println(qualy)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("name:", name, ok)

}
