package main

import (
	"fmt"
	"project-layout/filesHandling"
	"project-layout/safety"
)

func main() {
	l := filesHandling.ReadFile("advent1.txt")
	total := 0
	for _, v := range l {
		if safety.IsSafe(v) {
			total++
		}

	}
	fmt.Println(total)
}
