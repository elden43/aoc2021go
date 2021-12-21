package inputter

import (
	"bufio"
	"os"
)

// GetLines get slice of results by new line
func GetLines(fileName string) []string {
	input, _ := os.Open(fileName)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
