package syd

import (
	"fmt"
	"time"
)

func Sydlog(str string) {
	/* log */
	fmt.Print("\033[31m [DEBUG] \033[0m")
	fmt.Print("\033[33m ")
	fmt.Print(time.Now())
	fmt.Print(" \033[0m")
	fmt.Print("\t\t")
	fmt.Print("\033[34m ")
	fmt.Print(str)
	fmt.Println(" \033[0m")
}
