package main

import (
	"fmt"
)

func main() {
	//int数组
	//var arr [5]int
	//arr := [5]int{1, 2, 3, 4, 5}
	// arr := [5]bool{true, false}
	// fmt.Print(arr)
	//float 数组
	// slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
	// fmt.Println(len(slice2), cap(slice2)) // 3 5
	// slice2 = append(slice2, 1, 2, 3, 4)   // [0, 0, 0, 1, 2, 3, 4]
	// fmt.Println(len(slice2), cap(slice2)) // 7 12
	// // 子切片 [start, end)
	// sub1 := slice2[3:] // [1 2 3 4]
	// sub2 := slice2[:3] // [0 0 0]
	// //sub3 := slice2[1:4] // [0 0 1]
	// // 合并切片
	// combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]
	// fmt.Println(combined)

	//map
	// m1 := make(map[string]int) //map[key]value
	// m2 := map[string]string{
	// 	"Sam":   "Male",
	// 	"Alice": "Female",
	// }
	//m1["tom"]=18
	// fmt.Println(m1)
	// fmt.Println(m2)

	//pointer{
	// str := "golang"
	// var p *string = &str //&会获取变量的地址
	// fmt.Println(p) //*就是取内容，直接写是取地址
	// *p = "a"
	// fmt.Println(str)
	//}

	//参数的值传递{
	//如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值。
	//如果参数使用指针，对参数的传递将会影响到外部变量
	//
	// func add(num int) {
	// 	num += 1
	// }

	// func realAdd(num *int) {
	// 	*num += 1
	// }

	// func main() {
	// 	num := 100
	// 	add(num)
	// 	fmt.Println(num)  // 100，num 没有变化

	// 	realAdd(&num)
	// 	fmt.Println(num)  // 101，指针传递，num 被修改
	// }
	//}
	// 使用type自定义类型
	// 使用const初始化自定义的类型
	// 使用switch进行选择
	// type Gender int8
	// const (
	// 	MALE   Gender = 1
	// 	FAMALE Gender = 2
	// )

	// gender := FAMALE
	// fmt.Print(gender)

	// switch gender {
	// case FAMALE:
	// 	fmt.Println("female")
	// 	fmt.Print(gender)
	// case MALE:
	// 	fmt.Println("male")
	// default:
	// 	fmt.Println("unknown")
	// }
	//Go 语言的 switch 不需要 break，匹配到某个 case，执行完该 case 定义的行为后，默认不会继续往下执行。
	//如果需要继续往下执行，需要使用 fallthrough
}
