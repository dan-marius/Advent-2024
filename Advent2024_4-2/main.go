package main

import (
	"fmt"
	"project-layout/filesHandling"
)

var sum int64

func main() {

	// Read lines from the file
	lines := filesHandling.ReadFile("advent.txt")

	for l, line := range lines {
		//fmt.Println(l, " ", line)
		//fmt.Println("Sum: ", sum)
		for c, char := range line {
			//fmt.Printf("%d  %c\n", c, char)
			// Compare each character to A
			if char == 'A' {
				if l > 0 && l < len(line)-1 && c > 0 && c < len(lines)-1 {
					checkRight(l, c, &lines)
					checkLeft(l, c, &lines)
					checkUp(l, c, &lines)
					checkDown(l, c, &lines)
				}
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
func checkRight(line int, pos int, lines *[]string) {
	// Check if the position is not the last position
	if pos < len((*lines)[line])-1 {
		// Check if the character right is an M
		if (*lines)[line+1][pos+1] == 'M' {
			if (*lines)[line-1][pos+1] == 'M' {
				if (*lines)[line+1][pos-1] == 'S' {
					if (*lines)[line-1][pos-1] == 'S' {
						fmt.Println("Found a match (M Right) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkLeft(line int, pos int, lines *[]string) {
	// Check if the position is not the first position
	if pos > 0 {
		// Check if the character left is an M
		if (*lines)[line+1][pos-1] == 'M' {
			if (*lines)[line-1][pos-1] == 'M' {
				if (*lines)[line+1][pos+1] == 'S' {
					if (*lines)[line-1][pos+1] == 'S' {
						fmt.Println("Found a match (M Left) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkUp(line int, pos int, lines *[]string) {
	// Check if the line is not the first line
	if line > 0 {
		// Check if the character up is an M
		if (*lines)[line+1][pos-1] == 'M' {
			if (*lines)[line+1][pos+1] == 'M' {
				if (*lines)[line-1][pos-1] == 'S' {
					if (*lines)[line-1][pos+1] == 'S' {
						fmt.Println("Found a match (M Up) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkDown(line int, pos int, lines *[]string) {
	// Check if the line is not the last line
	if line < len(*lines)-1 {
		// Check if the character down is an M
		if (*lines)[line-1][pos-1] == 'M' {
			if (*lines)[line-1][pos+1] == 'M' {
				if (*lines)[line+1][pos-1] == 'S' {
					if (*lines)[line+1][pos+1] == 'S' {
						fmt.Println("Found a match (M Down) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
