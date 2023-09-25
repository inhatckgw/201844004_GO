package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	month := now.Month()
	fmt.Println(year, month)

	brokenGo := "G? r?cks!"
	replacerGo := strings.NewReplacer("?", "o")
	fixedGo := replacerGo.Replace(brokenGo)
	fmt.Println(fixedGo)

	fmt.Print("Input Score: ")
	reader := bufio.NewReader((os.Stdin))
	//input := reader.ReadString('\n') -> 복수의 변수를 갖는 ReadString
	//input, err := reader.ReadString('\n') -> 변수가 사용안됨
	//input, _ := reader.ReadString('\n') -> 사용할 변수가 한개인 경우 _(모든 에러 값 무시)으로 변수 선언 시 가능
	inputScore, err := reader.ReadString('\n')
	if err != nil { //conditional
		log.Fatal(err) //에러 보고 후 프로그램 종료
	}
	fmt.Println(inputScore)
}
