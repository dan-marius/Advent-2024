package main

import (
	"fmt"
	"project-layout/filesHandling"
	"sort"
)

func main() {
	// Example slice to demonstrate sorting
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]

	// Call the Hello function from the filesHandling package
	filesHandling.Hello()

	// Read numbers from the file
	a, b := filesHandling.ReadFile("advent1.txt")
	//fmt.Println(a, " then ", b)
	x := a[:]
	y := b[:]
	var sum uint64
	sort.Ints(x)
	sort.Ints(y)
	for _, v := range x {
		for _, z := range y {
			if v == z {
				sum += uint64(z)
			}
		}
	}
	fmt.Printf(" %v", sum)
}
