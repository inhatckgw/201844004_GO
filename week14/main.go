package main

import (
	"fmt"
)

func main() {
	var games map[int]string
	games = make(map[int]string)

	games[123] = "qwe"
	games[001] = "zxc"
	games[455] = "rty"
	games[124] = "fgh"
	games[354] = "vbn"
	games[294] = "uio"
	games[456] = "asd"

	for _, v := range games {
		fmt.Println(v)
	}

	games[001] = "plm"
	delete(games, 455)

	for k, v := range games {
		fmt.Println(k, v)
	}
}
