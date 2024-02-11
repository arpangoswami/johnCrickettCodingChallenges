package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// To deal with types cat <fileName> | ./ccwc -<flagName>
	if len(os.Args) == 2 {
		flagName := os.Args[1]
		validFlagName := false
		fileName := os.Stdin
		switch flagName {
		case "-l":
			validFlagName = true
			numLines := countNumberOfLines(fileName)
			fmt.Printf("%d\n", numLines)
		case "-c":
			validFlagName = true
			byteStream, _ := io.ReadAll(fileName)
			numberOfBytes := countNumberOfBytes(byteStream)
			fmt.Printf("%d\n", numberOfBytes)
		case "-w":
			validFlagName = true
			numWords := countNumberOfWords(fileName)
			fmt.Printf("%d\n", numWords)
		case "-m":
			validFlagName = true
			byteStream, _ := io.ReadAll(fileName)
			numberOfChars := countNumberOfChars(byteStream)
			fmt.Printf("%d\n", numberOfChars)
		}
		if validFlagName {
			return
		}
	}

	// To deal with types ./ccwc -<flagName> <fileName>
	var filenameNumberOfBytes, filenameNumberOfLines, fileNameNumberOfWords, fileNameNumberOfChars string

	flag.StringVar(&filenameNumberOfBytes, "c", "", "Count number of bytes for given filename (optional)")
	flag.StringVar(&filenameNumberOfLines, "l", "", "Count number of lines for given filename(optional)")
	flag.StringVar(&fileNameNumberOfWords, "w", "", "Count number of words for given filename(optional)")
	flag.StringVar(&fileNameNumberOfChars, "m", "", "Count number of words for given filename(optional)")
	flag.Parse()
	if fileNameNumberOfChars == "" && filenameNumberOfBytes == "" && filenameNumberOfLines == "" && fileNameNumberOfWords == "" {
		// When no flagName is provided, but by default we need numberOfLines, numberOfWords and numberOfBytes
		if len(os.Args) <= 1 {
			fmt.Println("Filename not provided, hence exiting")
			os.Exit(1)
		} else if len(os.Args) == 2 {
			fileName := os.Args[1]
			inputFile, err := os.Open(fileName)
			if err != nil {
				panic("Error occured during opening the given filename")
			}
			defer inputFile.Close()
			lineCount := countNumberOfLines(inputFile)
			_, _ = inputFile.Seek(0, 0) // Move the file ptr to the beginning of the file again
			wordCount := countNumberOfWords(inputFile)
			_, _ = inputFile.Seek(0, 0) // Move the file ptr to the beginning of the file again
			byteStream, _ := io.ReadAll(inputFile)
			byteCount := countNumberOfBytes(byteStream)
			fmt.Printf("%v\t%v\t%v\t%v\n", lineCount, wordCount, byteCount, fileName)
		} else {
			fmt.Println("Invalid number of cli arguments provided, exiting...")
			os.Exit(1)
		}
	} else if filenameNumberOfBytes != "" {
		byteStream, err := os.ReadFile(filenameNumberOfBytes)
		if err != nil {
			fmt.Println("Error reading given filename")
			os.Exit(1)
		}
		numBytes := countNumberOfBytes(byteStream)
		fmt.Printf("%d\t%v\n", numBytes, filenameNumberOfBytes)
	} else if fileNameNumberOfChars != "" {
		byteStream, err := os.ReadFile(fileNameNumberOfChars)
		if err != nil {
			panic(fmt.Sprintf("Error while opening file, invalid filename %v", fileNameNumberOfChars))
		}
		numChars := countNumberOfChars(byteStream)
		fmt.Printf("%d\t%v\n", numChars, fileNameNumberOfChars)
	} else if filenameNumberOfLines != "" {
		file, err := os.Open(filenameNumberOfLines)
		if err != nil {
			panic(fmt.Sprintf("Error while opening file, invalid filename %v", filenameNumberOfLines))
		}
		defer file.Close()
		numLines := countNumberOfLines(file)
		fmt.Printf("%d\t%v\n", numLines, filenameNumberOfLines)
	} else if fileNameNumberOfWords != "" {
		file, err := os.Open(fileNameNumberOfWords)
		if err != nil {
			panic(fmt.Sprintf("Error while opening file, invalid filename %v", fileNameNumberOfWords))
		}
		defer file.Close()
		numWords := countNumberOfWords(file)
		fmt.Printf("%d\t%v\n", numWords, fileNameNumberOfWords)
	}
}

func countNumberOfWords(input *os.File) int {
	scanner := bufio.NewScanner(input)
	words := 0
	for scanner.Scan() {
		partOfWord := false
		line := scanner.Text()
		for _, ch := range line {
			if unicode.IsSpace(ch) {
				partOfWord = false
			} else {
				if !partOfWord {
					words++
					partOfWord = true
				}
			}
		}
	}
	return words
}
func countNumberOfLines(input *os.File) int {
	scanner := bufio.NewScanner(input)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return lines
}
func countNumberOfChars(byteStream []byte) int {

	return utf8.RuneCountInString(string(byteStream))
}
func countNumberOfBytes(byteStream []byte) int {
	return len(byteStream)
}
