package main

import (
	"fmt"
)

func status(name string) {
	balls := map[string]int{"asd": 20, "qwe": 0}
	ball, exits := balls[name]
	if !exits {
		fmt.Println(name, "없음")
	} else if ball < 1 {
		fmt.Println(name, balls[name], "탈락")
	} else {
		fmt.Println(name, balls[name], "합격")
	}
}

func main() {
	// var games map[int]string
	// games = make(map[int]string)

	// games[123] = "qwe"
	// games[001] = "zxc"
	// games[455] = "rty"
	// games[124] = "fgh"
	// games[354] = "vbn"
	// games[294] = "uio"
	// games[456] = "asd"

	// for _, v := range games {
	// 	fmt.Println(v)
	// }

	// games[001] = "plm"
	// delete(games, 455)

	// for k, v := range games {
	// 	fmt.Println(k, v)
	// }

	//---------------------------------------------------------------------------//
	// games := map[int]string{456: "asd", 123: "sdw", 231: "zxc", 233: "fgh", 344: "sda", 001: "wer"}
	// for _, v := range games {
	// 	fmt.Println(v)
	// }

	// games[001] = "plm"
	// delete(games, 344)

	// for k, v := range games {
	// 	fmt.Println(k, v)
	// }
	status("asd")
	status("qwe")
	status("zxc")
}
