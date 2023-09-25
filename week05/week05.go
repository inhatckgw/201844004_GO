package main

import (
	"fmt"
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
}
