package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 20; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	//对数组
	nums := []int{1, 2, 3, 4, 5, 6}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	maps := map[string]int{
		"Tom": 18,
		"ann": 23,
	}
	for name, age := range maps {
		fmt.Println(name, age)
	}

}
