package filesHandling

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) []string {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return lines

}
