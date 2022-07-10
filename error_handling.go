package main

import (
	"fmt"
	//"os"
	//"errors"
)

// func main() {
// 	_, err := os.Open("filename.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

//******************************

// func hello(name string) error {
// 	if len(name) == 0 {
// 		return errors.New("error:name is null")
// 	}
// 	fmt.Println("hello", name)
// 	return nil
// }

// func main() {
// 	if err := hello(""); err != nil {
// 		fmt.Println(err)
// 	}
// }

//******************************

func get(index int) (ret int) {
	//类似于try catch，defer recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("some error happended!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	fmt.Println(get(2))
	fmt.Println("finished")
}
