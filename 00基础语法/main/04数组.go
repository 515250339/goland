/*
数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成，一旦声明了，数组的长度就固定了，不能动态变化。

`len()` 和 `cap()` 返回结果始终一样。
*/

// 声明数组 代码实例

package main

import (
	"fmt"
)

func main() {
	// 一维数组

	// 创建int类型的数组 长度为5 默认 value 都为0
	var arr_1 [5]int
	fmt.Println(arr_1)

	// 创建int类型的数组 长度为5 指定value
	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	// 同上
	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	// 创建int类型数组 长度为5 并指定下标的 value 未指定默认为0
	arr_4 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_4)

	// 创建int类型数组 不定长 给定6个 value
	arr_5 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_5)

	// 二维数组

	// 创建二维数组 有3组 每组5个元素 为 int 类型
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	// 同上
	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	//不定长 每组5个元素
	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	main_1()
	main_2()
}

// 注意事项

// 一、数组不可动态变化问题，一旦声明了，其长度就是固定的。

// var arr_1 = [5] int {1, 2, 3, 4, 5}
// arr_1[5] = 6
// fmt.Println(arr_1)

// 运行会报错：invalid array index 5 (out of bounds for 5-element array)

// 二、数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。

func main_1() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr)
	fmt.Println(arr)
}

func modifyArr(a [5]int) {
	a[1] = 20
	fmt.Println(a)
}

func main_2() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr_2(&arr)
	fmt.Println(arr)
}

func modifyArr_2(a *[5]int) {
	a[1] = 20
	fmt.Println(a)
}
