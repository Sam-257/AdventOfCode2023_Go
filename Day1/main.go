package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "./Day1/input3.txt"

var stringToNumberMap map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// Reads data from a file and returns it
func ReadFileData() string {
	textInput, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Unable to read data for %s due to %s\n", fileName, err)
	}
	return string(textInput)
}

func getCode(word string) (int, error) {
	firstDigit, lastDigit := "", ""
	wordLength := len(word)

	var setDigit = func (digit string) {
		if firstDigit == "" {
			firstDigit = digit
		} else {
			lastDigit = digit
		}
	}

	i := 0
	for i < wordLength {
		sliceBound := i + 5
		char := string(word[i])
		if sliceBound > wordLength {
			sliceBound = wordLength
		}
		isValid, digit := getStringNumber(word[i:sliceBound])
		// fmt.Println(isValid, digit, addend)
		if isValid {
			setDigit(digit)
		} else if char >= "0" && char <= "9" {
			setDigit(char)
		}
		i += 1
	}

	if lastDigit == "" {
		lastDigit = firstDigit
	}
	strCode := firstDigit + lastDigit
	return strconv.Atoi(strCode)
}

func getStringNumber(subWord string) (bool, string) {
	isValid, digit := false, ""
	for key, value := range stringToNumberMap {
		if strings.HasPrefix(subWord, key) {
			digit = value
			isValid = true
			break
		}
	}
	return isValid, digit
}

func main() {
	input := ReadFileData()
	words := strings.Split(input, "\n")
	var calibrationValue int

	for _, word := range words {
		code, _ := getCode(word)
		fmt.Println(code)
		calibrationValue += code
	}
	fmt.Println(calibrationValue)
}
