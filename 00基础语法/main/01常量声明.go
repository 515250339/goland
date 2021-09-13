/*
## 概述

在声明变量之前，咱们先了解下变量的数据类型，这篇文章主要涉及 字符串、布尔、数字，其他类型后面开篇再说。

## 数据类型

#### 字符串

`string`

只能用一对双引号（""）或反引号（``）括起来定义，不能用单引号（''）定义！

#### 布尔

`bool`

只有 true 和 false，默认为 false。

#### 数字

**整型**

`int8` `uint8`

`int16` `uint16`

`int32` `uint32`

`int64` `uint64`

`int` `uint`，具体长度取决于 CPU 位数。

**浮点型**

`float32` `float64`

## 常量声明

**常量**，在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

**单个常量声明**

第一种：const 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：const 变量名称 = 变量值

根据变量值，自行判断数据类型。

**多个常量声明**

第一种：const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：const 变量名称,变量名称 ...  = 变量值,变量值 ...
*/

// 代码实例
package main

import (
	"fmt"
)

func main() {
	// 创建常量 指定 string 类型
	const name string = "Tom"
	fmt.Println(name)

	// 创建常量 指定 int 类型
	const age = 30
	fmt.Println(age)

	// 创建多个常量 指定 string 类型
	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	// 创建多个常量 选择合适 value 的类型
	const name_3, age_1 = "TTTT", 50
	fmt.Println(name_3, age_1)

}
