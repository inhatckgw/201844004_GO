package main

import (
	"fmt"
	"os"
)

func main() {
	var a []string
	var b []bool

	b = append(b, true)
	fmt.Printf("%#v %#v\n", a, b) //zero value of empty slice

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(os.Args) //Args = Arguments
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[2]) //slice를 통으로 던질 때 slice명... 으로

}
