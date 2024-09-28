package count

import (
	"bufio"
	"fmt"
	"os"
)

func Lines(file *os.File, scanner *bufio.Scanner) (int, error) {
	// reset the file pointer to the start before reading lines
	file.Seek(0, 0)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file %s\n", err)
		return 0, err
	}

	return lineCount, nil
}
