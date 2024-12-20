package main

import (
	"fmt"
	"math"
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
	var difference int
	var total, distance float64
	sort.Ints(x)
	sort.Ints(y)
	fmt.Println(x, " then ", y)
	for i := 0; i < len(x); i++ {
		difference = x[i] - y[i]
		distance = math.Abs(float64(difference))
		total += distance
		fmt.Printf(" %d", int(total))
	}

}
