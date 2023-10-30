package main

import (
	"201844004_GO/week10/src/greeting"
	"201844004_GO/week10/src/mymath"
	"fmt"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Println(mymath.MyAbs(-7))
	fmt.Println(mymath.MyAbs(99))

	fmt.Println(mymath.MyPower(2, 10))
}
