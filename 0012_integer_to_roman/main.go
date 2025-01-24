package main

import "fmt"

var intToRomanMap = map[int]string{
	1000: "M",
	100:  "C",
	10:   "X",
	1:    "I",
}

var fivesIntToRomanMap = map[int]string{
	100: "D",
	10:  "L",
	1:   "V",
}

var foursIntToRomanMap = map[int]string{
	100: "CD",
	10:  "XL",
	1:   "IV",
}

var ninesIntToRomanMap = map[int]string{
	100: "CM",
	10:  "XC",
	1:   "IX",
}

var iterator = []int{1000, 100, 10, 1}

func intToRoman(num int) string {
	rest := num
	romanNumber := ""
	for j := 0; j < len(iterator); j += 1 {
		intDigit := iterator[j]
		numOfCurrentLevel := rest / intDigit
		if numOfCurrentLevel <= 0 {
			continue
		}

		if numOfCurrentLevel == 9 {
			romanNumber += ninesIntToRomanMap[intDigit]
			numOfCurrentLevel -= 9
		}

		if numOfCurrentLevel >= 5 {
			romanNumber += fivesIntToRomanMap[intDigit]
			numOfCurrentLevel -= 5
		}

		if numOfCurrentLevel >= 4 {
			romanNumber += foursIntToRomanMap[intDigit]
			numOfCurrentLevel -= 4
		}

		for i := 0; i < numOfCurrentLevel; i += 1 {
			romanNumber += intToRomanMap[intDigit]
		}

		rest %= intDigit
	}

	return romanNumber
}

func main() {
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(1))
	fmt.Println(intToRoman(10))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(500))
	fmt.Println(intToRoman(3999))
	fmt.Println(intToRoman(123))
	fmt.Println(intToRoman(458))
	fmt.Println(intToRoman(222))
	fmt.Println(intToRoman(1234))
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
