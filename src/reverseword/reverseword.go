package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e) // panic stops the ordinary flow of control and begins panicking
	}
}

func reverseWord(words []string) string {
	var buffer bytes.Buffer
	for itr := len(words) - 1; itr >= 0; itr-- {
		buffer.WriteString(words[itr] + " ")
	}
	return buffer.String()
}

func main() {

	fileHandler, e := os.Open("reversewordLarge.in")
	check(e)
	defer fileHandler.Close()

	fileScanner := bufio.NewScanner(fileHandler)
	fileScanner.Scan()
	numOfCases, err := strconv.Atoi(fileScanner.Text())
	check(err)

	for itr := 1; itr <= numOfCases; itr++ {
		fileScanner.Scan()
		fmt.Println("Case #", itr, "  ", reverseWord(strings.Split(fileScanner.Text(), " ")))
	}
}

// Input //
// 3
// this is a test
// foobar
// all your base

// Output //
// Case #1: test a is this
// Case #2: foobar
// Case #3: base your all
