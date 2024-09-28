package count

import (
	"bufio"
	"fmt"
	"os"
)

func Words(file *os.File, scanner *bufio.Scanner) (int, error) {
	// reset the file pointer to the start before reading the words
	file.Seek(0, 0)

	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file %s\n", err)
		return 0, err
	}

	return wordCount, nil
}
