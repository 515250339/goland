package main

import (
	"fmt"
)

// 循环 array
func main() {
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d  array=%v \n", len(person), cap(person), person)

	fmt.Println("")

	for k, v := range person {
		fmt.Printf("person[%d]:%s \n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]: %s \n", i, person[i])
	}

	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s \n", i, person[i])
	}

	// 使用空白符号
	for _, name := range person {
		fmt.Println("name=", name)
	}
	main2()
	main3()
}

// 循环 slice
func main2() {
	fmt.Println("---------slice-----------")
	person := []string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d slice=%v \n", len(person), cap(person), person)

	fmt.Println("")

	for k, v := range person {
		fmt.Printf("person[%d] : %s \n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]: %s \n", i, person[i])
	}

	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s \n", i, person[i])
	}

	fmt.Println("")
	// 使用空白符
	for _, name := range person {
		fmt.Println("name=", name)
	}
}

// 循环map
func main3() {
	fmt.Println("")
	fmt.Println("---------map-----------")

	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}

	fmt.Println("len=%d map=%v \n", len(person), person)

	fmt.Println("")

	for k, v := range person {
		fmt.Printf("person[%d]: %s \n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]=%s \n", i, person[i])
	}

	fmt.Println("")

	for i := 1; i <= len(person); i++ {
		fmt.Printf("person[%d]=%s \n", i, person[i])
	}

	fmt.Println("")

	for _, name := range person {
		fmt.Println("name = ", name)
	}
	main4()
}

// break 跳出当前循环，可⽤于 for、switch、select。
func main4() {
	fmt.Println("")
	fmt.Println("---------break-----------")

	for i := 1; 1 < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	main5()

}

// continue 跳过本次循环，只能用于 for。
func main5() {
	fmt.Println("")
	fmt.Println("---------continue-----------")

	for i := 1; i < 10; i++ {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
	main6()
}

// goto 改变函数内代码执行顺序，不能跨函数使用。
func main6() {
	fmt.Println("")
	fmt.Println("---------goto-----------")

	for i := 1; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println(i)
	}

END:
	fmt.Println("end")

	main7()
}

// switch 相当于 if elif
func main7() {
	fmt.Println("")
	fmt.Println("---------switch-----------")

	for i := 1; i <= 7; i++ {

		fmt.Printf("当 i = %d 时：\n", i)

		switch i {
		case 1:
			fmt.Println("输出 i=", i)
		case 2:
			fmt.Println("输出 i=", i)
		case 3:
			fmt.Println("输出 i=", i)
		case 4, 5, 6:
			fmt.Println("输出 i=", i)
		default:
			fmt.Println("输出 i", "默认")
		}
	}
}
