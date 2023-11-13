package main

import "fmt"

func main() {
	//var s []int
	//s = make([]int, 5)

	s := [5]int{1, 2, 3, 4, 5}

	s[0] = 99
	s[4] = 11

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, value1 := range s {
		fmt.Println(value1)
	}

	//slice1 := s[0:4]
	// for _, value2 := range slice1 {
	// 	fmt.Println(value2)
	// }
}
