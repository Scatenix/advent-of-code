package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// FileExists checks if file exists on the file system
func FileExists(file string) error {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File: %s does not exist\n", file)
		return err
	}
	return nil
}

// ReadPuzzleFile Read a file by creating a bufio Scanner.
// Provide a function f to handle each line read with scanner.Text() by adding it to collection[T]
// and returning it for the next scan.
//
// example lineHandler:
//	lineHandler := func(line string, ret string) string {
//			ret += line
//			return ret
//		}
//
// File stream will be closed after execution
//
// Returns collection[T] and scanner.Err()
func ReadPuzzleFile[T any](file *os.File, f func(line string, col *T) *T, col *T) error {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		col = f(scanner.Text(), col)
	}

	defer closeFile(file)
	return scanner.Err()
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while closing file %s:\n%v\n",f.Name(), err)
		os.Exit(1)
	}
}