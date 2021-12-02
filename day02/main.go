package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processPart1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	depth := 0
	moves := 0
	for scanner.Scan() {
		s := scanner.Text()
		result := strings.Split(s, " ")
		if result[0] == "forward" {
			if i, err := strconv.Atoi(result[1]); err == nil {
				moves += i
			}
		} else {
			if i, err := strconv.Atoi(result[1]); err == nil {
				if result[0] == "up" {
					depth -= i
				} else {
					depth += i
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return (moves * depth)
}

func processPart2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	depth := 0
	moves := 0
	aim := 0
	for scanner.Scan() {
		s := scanner.Text()
		result := strings.Split(s, " ")
		if result[0] == "forward" {
			if i, err := strconv.Atoi(result[1]); err == nil {
				moves += i
				depth = depth + (aim * i)
			}
		} else {
			if i, err := strconv.Atoi(result[1]); err == nil {
				if result[0] == "up" {
					aim -= i
				} else {
					aim += i
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return (moves * depth)
}

func main() {
	fmt.Println("Part 1")
	fmt.Println(processPart1("input"))
	fmt.Println("Part 2")
	fmt.Println(processPart2("input"))
}
