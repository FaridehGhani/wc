package count

import (
	"fmt"

	"io/ioutil"
)

func Chars(filename string) (int, error) {
	// read the entire file into a string
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("err reading file: %s\n", err)
		return 0, err
	}

	// convert the file content to a string and count the number of runes (characters)
	return len([]rune(string(content))), nil
}
