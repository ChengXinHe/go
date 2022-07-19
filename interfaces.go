package main

import (
	"fmt"
)

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string { //go 中不需要显式声明实现了哪些接口，只需要实现对应的方法
	return stu.name //检查有没有全部实现方法，将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student{ //实例化student后 强制转化为person
		name: "Tom",
		age:  18,
	}

	fmt.Println(p.getName()) // Tom
}

//空接口可以表示任意类型
// func main() {
// 	m := make(map[string]interface{})
// 	m["age"] = 18
// }
