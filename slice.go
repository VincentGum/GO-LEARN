package main
import (
	"fmt"
)

// 切片是对数组的抽象
// 切片可以动态改变长度，数组不可以

func newSlice() {
	// 声明一个切片

	var a []int = make([]int, 10)
	a = append(a, 1)
	fmt.Println(a)

	b := make([]string, 3)
	var i int
	for i=0; i<3; i++ {
		b[i] = "a"
	}
	fmt.Println(b)
	b = append(b, "b")
	fmt.Println(b)

	c := make([]int, 0, 6)
	fmt.Println(c)
	for i=0; i<6; i++ {
		c = append(c, i)
	}
	fmt.Println(c)
	fmt.Println(cap(c))
	fmt.Println(len(c))
}