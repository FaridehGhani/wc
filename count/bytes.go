package count

import (
	"fmt"
	"os"
)

func Bytes(file *os.File) (int, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return 0, err
	}

	return int(fileInfo.Size()), nil
}
