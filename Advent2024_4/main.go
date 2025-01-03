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
			// Compare each character to X
			if char == 'X' {
				checkLeftUp(l, c, &lines)
				checkLeftDown(l, c, &lines)
				checkRightUp(l, c, &lines)
				checkRightDown(l, c, &lines)
				checkRight(l, c, &lines)
				checkLeft(l, c, &lines)
				checkUp(l, c, &lines)
				checkDown(l, c, &lines)
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
func checkLeftUp(line int, pos int, lines *[]string) {
	// Check if the line is not the first line
	if line > 2 {
		// Check if the position is not the first position
		if pos > 2 {
			// Check if the character top left is an M
			if (*lines)[line-1][pos-1] == 'M' {
				if (*lines)[line-2][pos-2] == 'A' {
					if (*lines)[line-3][pos-3] == 'S' {
						fmt.Println("Found a match (Left Up) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkLeftDown(line int, pos int, lines *[]string) {
	// Check if the line is not the last line
	if line < len(*lines)-3 {
		// Check if the position is not the first position
		if pos > 2 {
			// Check if the character bottom left is an M
			if (*lines)[line+1][pos-1] == 'M' {
				if (*lines)[line+2][pos-2] == 'A' {
					if (*lines)[line+3][pos-3] == 'S' {
						fmt.Println("Found a match (Left Down) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkRightUp(line int, pos int, lines *[]string) {
	// Check if the line is not the first line
	if line > 2 {
		// Check if the position is not the last position
		if pos < len((*lines)[line])-3 {
			// Check if the character top right is an M
			if (*lines)[line-1][pos+1] == 'M' {
				if (*lines)[line-2][pos+2] == 'A' {
					if (*lines)[line-3][pos+3] == 'S' {
						fmt.Println("Found a match (Right Up) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkRightDown(line int, pos int, lines *[]string) {
	// Check if the line is not the last line
	if line < len(*lines)-3 {
		// Check if the position is not the last position
		if pos < len((*lines)[line])-3 {
			// Check if the character bottom right is an M
			if (*lines)[line+1][pos+1] == 'M' {
				if (*lines)[line+2][pos+2] == 'A' {
					if (*lines)[line+3][pos+3] == 'S' {
						fmt.Println("Found a match (Right Down) at line ", line, " pos ", pos)
						sum++
					}
				}
			}
		}
	}
}
func checkRight(line int, pos int, lines *[]string) {
	// Check if the position is not the last position
	if pos < len((*lines)[line])-3 {
		// Check if the character right is an M
		if (*lines)[line][pos+1] == 'M' {
			if (*lines)[line][pos+2] == 'A' {
				if (*lines)[line][pos+3] == 'S' {
					fmt.Println("Found a match (Right) at line ", line, " pos ", pos)
					sum++
				}
			}
		}
	}
}
func checkLeft(line int, pos int, lines *[]string) {
	// Check if the position is not the first position
	if pos > 2 {
		// Check if the character left is an M
		if (*lines)[line][pos-1] == 'M' {
			if (*lines)[line][pos-2] == 'A' {
				if (*lines)[line][pos-3] == 'S' {
					fmt.Println("Found a match (Left) at line ", line, " pos ", pos)
					sum++
				}
			}
		}
	}
}
func checkUp(line int, pos int, lines *[]string) {
	// Check if the line is not the first line
	if line > 2 {
		// Check if the character up is an M
		if (*lines)[line-1][pos] == 'M' {
			if (*lines)[line-2][pos] == 'A' {
				if (*lines)[line-3][pos] == 'S' {
					fmt.Println("Found a match (Up) at line ", line, " pos ", pos)
					sum++
				}
			}
		}
	}
}
func checkDown(line int, pos int, lines *[]string) {
	// Check if the line is not the last line
	if line < len(*lines)-3 {
		// Check if the character down is an M
		if (*lines)[line+1][pos] == 'M' {
			if (*lines)[line+2][pos] == 'A' {
				if (*lines)[line+3][pos] == 'S' {
					fmt.Println("Found a match (Down) at line ", line, " pos ", pos)
					sum++
				}
			}
		}
	}
}
