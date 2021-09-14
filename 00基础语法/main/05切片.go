/*
## 概述

切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

`len()` 和 `cap()` 返回结果可相同和不同。

new 的作用是初始化一个指向类型的指针(*T)，make 的作用是为 slice，map 或 chan 初始化并返回引用(T)。

切片可以通过 len() 方法获取长度，可以通过 cap() 方法获取容量。
len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

*/

// 声明切片 代码实例

package main

import (
	"fmt"
)

func main() {

	fmt.Println("----------声明切片-------")
	var sli_1 []int // nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = []int{} // 空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_2), cap(sli_2), sli_2)

	// 俩种创建方式
	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)

	// 使用 make 创建
	var sli_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)

	main_2()
}

// 截取切片

func main_2() {
	fmt.Println("----------截取切片-------")
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	// 包头不包尾
	fmt.Println("sli[0]==", sli[0])
	fmt.Println("sli[:]==", sli[:])
	fmt.Println("sli[1:]==", sli[1:])
	fmt.Println("sli[:4]==", sli[:4])

	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	fmt.Println("sli[:3:5]==", sli[:3:5])
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli[0:3:4]), cap(sli[0:3:6]), sli[0:3:4])

	main_3()
}

// 追加切片
func main_3() {
	fmt.Println("----------追加切片-------")
	// sli := make([] int, 6, 9)
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	sli = append(sli, 8)
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	sli = append(sli, 9)
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	sli = append(sli, 10)
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	main_4()
}

// 删除切片
func main_4() {
	fmt.Println("-------删除切片-------")
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)
	// 删除尾部俩个元素
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli[0:len(sli)-2]), cap(sli[0:len(sli)-2]), sli[0:len(sli)-2])
	// 删除头部俩个数据
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli[2:]), cap(sli[2:]), sli[2:])
	// 删除中间2个
	sli = append(sli[:len(sli)/2], sli[len(sli)/2+2:]...)
	fmt.Printf("len=%d, cap=%d, sli=%v\n", len(sli), cap(sli), sli)

}
