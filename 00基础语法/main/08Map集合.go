/*
Map 集合是无序的 key-value 数据结构。

Map 集合中的 key / value 可以是任意类型，但所有的 key 必须属于同一数据类型，所有的 value 必须属于同一数据类型，key 和 value 的数据类型可以不相同。
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 新建map对象
	var p1 map[int]string
	// 初始化map 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	p1 = make(map[int]string)
	// 存储 key(int) value(string)
	p1[1] = "Tom"
	fmt.Println("p1=", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2=", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3=", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4=", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5=", p5)

	p6 := map[int]string{
		1: "Tom",
	}
	p6[1] = "Tom"
	fmt.Println("p6=", p6)
	json_main()
}

func json_main() {

	fmt.Println("-------序列化-------")
	res := make(map[int]string)
	res[200] = "ok"
	fmt.Println(res)

	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("errs", errs)
	}
	fmt.Println("json data ：", string(jsons))

	// 反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("errs", errs)
	}
	fmt.Println("res2=", res2)
	main_2()
}

func main_2() {
	fmt.Println("-------编辑和删除-------")
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Println("data: ", person)
	// 删除
	delete(person, 2)
	fmt.Println("delete: ", person)
	// 添加
	person[2] = "jack"
	// 修改
	person[3] = "wang"
	fmt.Println("data: ", person)

}
