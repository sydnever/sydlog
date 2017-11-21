package syd

import (
	"fmt"
	"time"
)

func Sydlog(str string) {
	fmt.Print("[DEBUG]")
	fmt.Print(time.Now())
	fmt.Print("  ")
	fmt.Println(str)
}
