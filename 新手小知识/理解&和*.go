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
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	main2()
}

func main2() {
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
