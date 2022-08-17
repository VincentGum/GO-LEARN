package main

import (
	"fmt"
)

// struct用作结构体
// 定义结构体的方式

type People struct {
	firstName	string
	lastName	string
	age			int
	gender 		string
	pass 		bool
}

func newStruct() {
	var john = People{"john", "muston", 34, "male", true}
	fmt.Println(john)
	fmt.Println(john.firstName)
}

func getStuInfo(stu People) string {
	info := "This People is %s.%s, aged %d"
	result := fmt.Sprintf(info, stu.firstName, stu.lastName, stu.age)
	return result
}