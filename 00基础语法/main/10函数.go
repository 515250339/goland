/*
值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递：引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
*/

package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// 参数
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
*/

// 转换为json格式
func toJson(res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data: ", string(jsons))
}
// 修改数据
func setData(res *Result) {
	res.Code = 500
	res.Message = "no"

}


func main() {
	var res Result
	res.Code = 200
	res.Message = "ok"
	/*
		调用 函数
	   	&res 指向 rea 指针，res 变量的地址
	*/
	toJson(&res)
	setData(&res)
	toJson(&res)
	var a int = 200
	var b int = 300
	var ret int
	var msg string
	ret, msg = max(a, b)
	fmt.Println(ret, msg)
	fmt.Println(square(9))
	var nexNumber = getSequence()
	fmt.Println(nexNumber())
	fmt.Println(nexNumber())
	fmt.Println(nexNumber())
	fmt.Println(nexNumber())
	fmt.Println(nexNumber())
}

// 寻找两数最大值，返回多参数
func max(a, b int) (int, string) {
	var res int
	if a > b {
		res = a
	} else {
		res = b
	}
	return res, "ok"

}

// Go 语言可以很灵活的创建函数，并作为另外一个函数的实参
func square(s float64) float64 {
	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
	return getSquareRoot(s)
	

}

// 闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
