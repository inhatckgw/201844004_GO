package main

import "fmt"

func main() {
	//var s []int
	//s = make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	as := a[0:2]
	as[1] = 8
	fmt.Println(a)

	b := []int{4, 3, 2, 1}
	bs := b[0:]
	bs[3] = 15
	fmt.Println(b)

	c := make([]string, 4, 5) //length(4) != capacity(5)
	c[0] = "a"
	c[2] = "b"
	c[3] = "c"
	cs := c[0:2]
	cs[1] = "z"
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &c[0], &cs[0], &c[1])

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }

	// for _, value1 := range s {
	// 	fmt.Print(value1, " ")
	// }
	// fmt.Println(" ")

	// slice1 := s[0:4]
	// for _, value2 := range slice1 {
	//  	fmt.Println(value2)
	// }

	// test := [3]string{"a", "b", "c"}
	// testS := test[0:2]
	// fmt.Println(len(testS))

	// testS2 := test[0:]
	// fmt.Println(testS2[1])

	d1 := []string{"d1", "d1"}
	d2 := append(d1, "d2", "d2")
	d3 := append(d2, "d3", "d3")
	d4 := append(d3, "d4", "d4")

	fmt.Println(d1, d2, d3, d4)
	d4[0] = "d9"
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(d1, len(d1), cap(d1))
	fmt.Println(d4, len(d4), cap(d4))

}
