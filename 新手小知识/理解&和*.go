/*
&是“取地址运算符”，是从一个变量获取地址
*是“解引用运算符”，可以简单理解为“从地址取值”， 是&的逆运算
你 testd 是一个 Test*类型，也就是指向 Test 的指针
然后&testd 就是 testd 变量本身的地址，类型应该是 Test 的指针的指针
*/

package main

import "fmt"

type Test struct {
	name string
}

func main() {
	testa := Test{"test"}
	fmt.Println(testa)
	//结果{test}

	testb := &Test{"test"}
	fmt.Println(testb)
	//结果 &{test}

	testc := &Test{"test"}
	fmt.Println(*testc)
	//结果 {test}

	testd := &Test{"test"}
	fmt.Println(&testd)
	//结果 0xc000006030

	var a int = 1
	fmt.Println(a)
	//结果 1
	fmt.Println(&a)
	//结果 0xc00000c0d8
}
