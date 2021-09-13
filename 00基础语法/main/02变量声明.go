/*
## 变量声明

**单个变量声明**

第一种：var 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：var 变量名称 = 变量值

根据变量值，自行判断数据类型。

第三种：变量名称 := 变量值

省略了 var 和数据类型，变量名称一定要是未声明过的。

**多个变量声明**

第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：var 变量名称,变量名称 ...  = 变量值,变量值 ...

第三种：变量名称,变量名称 ... := 变量值,变量值 ...
*/

// 代码实例
package main

import (
	"fmt"
)

func main() {
	// 创建变量 指定类型 uint8
	var age_1 uint8 = 31

	// 创建变量 自行判断数据类型
	var age_2 = 32

	// 省略了 var 和数据类型，变量名称一定要是未声明过的。
	age_3 := 33
	fmt.Println(age_1, age_2, age_3)

	// 创建多个变量 指定为 int 类型
	var age_4, age_5, age_6 int = 31, 32, 33
	fmt.Println(age_4, age_5, age_6)

	// 创建多个变量 自行判断数据类型
	var name_1, age_7 = "Tom", 33
	fmt.Println(name_1, age_7)

	// 创建多个变量 自行判断数据类型 省略了 var 和数据类型，变量名称一定要是未声明过的。
	name_2, is_boy, height := "Jay", true, 180.66
	fmt.Println(name_2, is_boy, height)

}
