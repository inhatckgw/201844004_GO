package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	var a int = 6
	var b = 2.7
	var c, d string
	var e float32 = 3.14
	var f string
	var g bool
	h := 'Z'
	i := "문자열"
	J := "변수명이 대문자로 시작하면 다른 패키지에서 변수를 사용할 수 있다"
	koreanZombie := "정찬성" //Go 커뮤니티에 맞는 변수명은 뒷 단어는 대문자로

	//a = 4
	//b = 2.71
	c = "abc"
	d = "Go..."

	fmt.Println(float64(a) > b) //형변환
	fmt.Println(a * int(b))

	fmt.Println(a, b, c, d, e, f, g, h, i, J, koreanZombie) //값 출력
	fmt.Println(reflect.TypeOf((b)))                        //float64
	fmt.Println(reflect.TypeOf((h)))                        //int32
	fmt.Println(reflect.TypeOf((i)))                        //string

	fmt.Println(reflect.TypeOf((42)))           //int
	fmt.Println(reflect.TypeOf((3.1415)))       //float64
	fmt.Println(reflect.TypeOf((true)))         //bool
	fmt.Println(reflect.TypeOf(("Hello, Go!"))) //string

	fmt.Println('A')

	fmt.Println(math.Floor(2.753)) //내림
	fmt.Println(math.Ceil(2.35))   //올림

	fmt.Println(strings.Title("오픈소스\t 프로그래밍"))
	fmt.Println(strings.Title("head first go")) //Title은 영문자의 첫 글자를 무조건 대문자로
}
