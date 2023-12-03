package main

import (
	"aoc/2023/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := util.FileReader("01", "2023")
	fmt.Println("Part 1: ")
	fmt.Println(part1(input))
	fmt.Println("--------")
	fmt.Println("Part 2: ")
	fmt.Println(part2(input))
}

func part1(input []string) string {
	var result int = 0

	for _, line := range input {
		d := ""

		digits := regexp.MustCompile(`\d`).FindAllString(line, -1)
		d += digits[0]
		d += digits[len(digits)-1:][0]
		calibration_value, _ := strconv.Atoi(d)
		result += calibration_value
	}
	return strconv.Itoa(result)
}

func part2(input []string) string {
	var result int = 0

	for _, line := range input {
		digits := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`).FindAllString(line, -1)
		digit_map := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
			"1":     1,
			"2":     2,
			"3":     3,
			"4":     4,
			"5":     5,
			"6":     6,
			"7":     7,
			"8":     8,
			"9":     9,
		}
		tens, _ := digit_map[digits[0]]

		reverse_digit_map := map[string]int{
			"eno":   1,
			"owt":   2,
			"eerht": 3,
			"ruof":  4,
			"evif":  5,
			"xis":   6,
			"neves": 7,
			"thgie": 8,
			"enin":  9,
			"1":     1,
			"2":     2,
			"3":     3,
			"4":     4,
			"5":     5,
			"6":     6,
			"7":     7,
			"8":     8,
			"9":     9,
		}

		reverse_digits := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`).FindAllString(reverse(line), -1)

		ones, _ := reverse_digit_map[reverse_digits[0]]

		fmt.Println(line)
		fmt.Println(10*tens + ones)
		result += 10*tens + ones
	}
	return strconv.Itoa(result)
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func notes() {
	// strconv.Itoa - convert int to string
	// strconv.Atoi - convert string to int
	// slice = append(slice, value)
	// sort.Slice(a, func(i, j int) bool {	return a[i] < a[j] })
}
