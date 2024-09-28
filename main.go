package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"wc/count"
)

func main() {
	// define a flag for the -c, -l, -w and -m option
	countBytes := flag.Bool("c", false, "count the number of bytes in a file")
	countLines := flag.Bool("l", false, "count the number of lines in a file")
	countWords := flag.Bool("w", false, "count the number of words in a file")
	countChars := flag.Bool("m", false, "count the number of characters in a file")
	flag.Parse()

	// ensure file is provided
	if len(flag.Args()) < 1 {
		fmt.Println("usage: ccwc [-c | -l | -w | -m] <filename>")
		return
	}

	// get the filename from the arguments
	filename := flag.Args()[0]

	// open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var byteCount, lineCount, wordCount, charCount int

	// if no flags are provided, default to -c, -l and -w behavior
	if !*countBytes && !*countLines && !*countWords && !*countChars {
		*countBytes, *countLines, *countWords = true, true, true
	}

	// if the flag -c is passed, count bytes
	if *countBytes {
		if byteCount, err = count.Bytes(file); err != nil {
			return
		}

	}

	// if the -w flag is passed, count words
	if *countWords {
		if wordCount, err = count.Words(file, scanner); err != nil {
			return
		}
	}

	// if the -l flag is passed, count lines
	if *countLines {
		if lineCount, err = count.Lines(file, scanner); err != nil {
			return
		}
	}

	// if the -m flag is passed, count characters (runes)
	if *countChars {
		if charCount, err = count.Chars(filename); err != nil {
			return
		}
		fmt.Printf("%8d %s\n", charCount, filename)

		return
	}

	// Print the output in the format: <line count> <word count> <byte count> <filename>
	if *countLines && *countWords && *countBytes {
		fmt.Printf("%8d %8d %8d %s\n", lineCount, wordCount, byteCount, filename)
	} else {
		// Print results for individual flags if provided
		if *countLines {
			fmt.Printf("%d ", lineCount)
		}
		if *countWords {
			fmt.Printf("%d ", wordCount)
		}
		if *countBytes {
			fmt.Printf("%d ", byteCount)
		}
		fmt.Printf("%s\n", filename)
	}
}
