package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func main() {
	//字段不一定都要赋值，可以被赋予默认值
	stu := &Student{
		name: "Tom",
		age:  18,
	}
	msg := stu.hello("JAck")
	fmt.Println(msg)
	//stu是实例名，*Student是类型。可以通过实例名访问字段name和方法
	stu2 := new(Student)
	fmt.Println(stu2.hello("alice"))
}
