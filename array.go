package main

import (
	"fmt"
	"reflect"
)

func oneDimArray() {
	// 最基础的初始化方式
	var a = []int{1,2,3,4}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
	a = append(a, 1)
	fmt.Println(a)

	// 给定初始长度
	var b = [10]int{1,2,3,4}
	b[5] = 9
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)

	// 快速声明列表
	c := [3]int{1,2,3}
	fmt.Println(c)

	// 对于长度不确定的数组
	d := [...]int{1,1,1,1,1}
	fmt.Println(d)
	
	var i int
	for i=0; i < len(d); i++ {
		fmt.Println(d[i])
	}
}

func multiDimArray() {
	var a = [][]int{}
	
	l1 := []int{1,2,3}
	l2 := []int{1}
	l3 := []int{1,2,3,5}

	a = append(a, l1)
	a = append(a, l2)
	a = append(a, l3)

	fmt.Println(a)
}