package filesHandling

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Hello() {
	fmt.Println("Hello from openRead")
}

func ReadFile(filePath string) []string {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	defer fd.Close()
	var lines []string
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
