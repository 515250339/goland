/*
结构体是将零个或多个任意类型的变量，组合在一起的聚合数据类型，也可以看做是数据的集合。

*/

package main

import (
	"fmt"
)

// 数据集
type Person struct {
	name string
	age  int
}

func main() {
	// 使用方式
	var p1 Person
	p1.name = "张三"
	p1.age = 30
	fmt.Printf("姓名：%s, 年龄：%d\n", p1.name, p1.age)

	var p2 = Person{name: "李四", age: 22}
	fmt.Printf("姓名：%s, 年龄：%d\n", p2.name, p2.age)

	p3 := Person{name: "李四", age: 22}
	fmt.Printf("姓名：%s, 年龄：%d\n", p3.name, p3.age)

	// 匿名结构体
	p4 := struct {
		name string
		age  int
	}{name: "匿名", age: 998}
	fmt.Println("p4=", p4)

}
