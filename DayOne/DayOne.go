package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func main() {
    file, err := os.Open("FileToParse.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    totalCounter := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		regexToUse := regexp.MustCompile("[0-9]")
		intsOnly := regexToUse.FindAllString(scanner.Text(), -1)
		numberOfInts := (len(intsOnly) - 1)
		
		combinedNumber, err := strconv.Atoi(intsOnly[0] + intsOnly[numberOfInts])
		if err != nil {
			log.Fatal(err)
		}
		totalCounter = totalCounter + combinedNumber
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("The answer is...")
    fmt.Println(totalCounter)
}