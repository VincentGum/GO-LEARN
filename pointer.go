package main

import (
	"fmt"
)

func task() {
	fmt.Println("hahahaha")
}

func taskAssign() {
	fmt.Println("\ntaskAssign")
	var a int = 111
	fmt.Println(a)  // 打印变量的地址
	fmt.Println(&a) // 打印变量的地址

	var b *int // 定义指针变量
	b = &a     // 指针变量赋值
	fmt.Println("\n b:")
	fmt.Println(b)  // b中保存的地址，即a的地址
	fmt.Println(&b) // b自身的地址
	fmt.Println(*b) // 变量的值

	var c **int
	c = &b
	fmt.Println("\n c:")
	fmt.Println(c)  // b中保存的地址，即a的地址
	fmt.Println(&c) // b自身的地址
	fmt.Println(*c) // 变量的值
}

func swapByVal(a int, b int) {
	temp := a
	a = b
	b = temp
}

func swapByRef(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
