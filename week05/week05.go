package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

	// if score >= 90 {
	// 	grade := "A Grade"
	// } else {
	// 	grade := "Under A Grade"
	// }
	// fmt.Println("You Will Get:", grade) -> undefined: grade

	inputScore = strings.TrimSpace(inputScore)       //문자열 내의 공백 제거
	score, err := strconv.ParseFloat(inputScore, 32) // 문자열을 float, int로 변환
	var grade string                                 //if 문 밖에 변수를 선언 해줌으로써 if문 내의 값을 밖에서도 사용
	if score >= 90 {                                 //string이랑 int 비교 에러발생
		grade = "A Grade"
	} else {
		grade = "Under A Grade"
	}
	//fmt.Println(inputScore)
	fmt.Println("Your Score: ", inputScore, "You Will Get: ", grade)

	targetRandom := rand.Intn(6) + 1 //0-5 사이의 정수
	fmt.Println(targetRandom)

	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1
	fmt.Println("Game Number (1 - 100)")
	fmt.Println(answer)

	readerGame := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You have: ", 10-i, "Chance")
		fmt.Print("Game Num:")
		input, err := readerGame.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if guess < answer {
			fmt.Println("Your Guess is lower than answer")
		} else if guess > answer {
			fmt.Println("Your Guess is higher than answer")
		} else if guess == answer {
			fmt.Println("Answer")
			break
		}
	}

}
