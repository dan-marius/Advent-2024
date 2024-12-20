package filesHandling

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Hello() {
	fmt.Println("Hello from openRead")
}

func ReadFile(filePath string) ([1000]int, [1000]int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	defer fd.Close()
	i := 0
	var x, y [1000]int
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		//fmt.Println(i)
		a, b := splitLineToInts(scanner.Text())
		//fmt.Println(a, b)
		x[i] = a
		y[i] = b
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return x, y
}

func splitLineToInts(line string) (int, int) {
	fields := strings.Fields(line) // Split the string by whitespace

	first, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0
	}

	second, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, 0
	}

	return first, second
}
