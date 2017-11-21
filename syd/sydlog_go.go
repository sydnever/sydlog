package syd

import (
	"fmt"
	"time"
)

func Sydlog(str string) {
	fmt.Print(time.Now())
	fmt.Println(str)
}
