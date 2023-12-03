package main

import (
	"aoc/2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.FileReader("02", "2023")
	fmt.Println("Part 1: ")
	fmt.Println(part1(input))
	fmt.Println("--------")
	fmt.Println("Part 2: ")
	fmt.Println(part2(input))
}

type Game struct {
	red   int
	green int
	blue  int
}

func part1(input []string) string {
	result := 0
	target_game := Game{12, 13, 14}
	for idx, line := range input {
		games := build_games_from_string(line)
		least_possible_game := determine_least_possible_game(games)
		if least_possible_game.red <= target_game.red && least_possible_game.blue <= target_game.blue && least_possible_game.green <= target_game.green {
			result += idx + 1
		}
	}
	return strconv.Itoa(result)
}

func part2(input []string) string {
	result := 0
	for _, line := range input {
		games := build_games_from_string(line)
		least_possible_game := determine_least_possible_game(games)
		result += least_possible_game.red * least_possible_game.blue * least_possible_game.green
	}
	return strconv.Itoa(result)
}

func build_game_from_string(game_string string) Game {
	result := Game{0, 0, 0}
	for _, ball_string := range strings.Split(game_string, ", ") {
		value, _ := strconv.Atoi(strings.Split(ball_string, " ")[0])
		color := strings.Split(ball_string, " ")[1]
		switch color {
		case "red":
			result.red += value
		case "blue":
			result.blue += value
		case "green":
			result.green += value
		}
	}
	return result
}

func build_games_from_string(line string) []Game {
	result := []Game{}
	games := strings.Split(line, ": ")[1]
	for _, game_string := range strings.Split(games, "; ") {
		result = append(result, build_game_from_string(game_string))
	}
	return result
}

func determine_least_possible_game(games []Game) Game {
	result := Game{0, 0, 0}
	for _, game := range games {
		if game.red > result.red {
			result.red = game.red
		}
		if game.blue > result.blue {
			result.blue = game.blue
		}
		if game.green > result.green {
			result.green = game.green
		}
	}

	return result
}

func notes() {
	// strconv.Itoa - convert int to string
	// strconv.Atoi - convert string to int
	// slice = append(slice, value)
	// sort.Slice(a, func(i, j int) bool {	return a[i] < a[j] })
}
