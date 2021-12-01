package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Increased(numbers []int) int {
	increased := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increased++
		}
	}
	return increased
}

func SlidingWindows(numbers []int) int {
	increased := 0
	for i := 3; i < len(numbers); i++ {
		sum1 := numbers[i-1] + numbers[i-2] + numbers[i-3]
		sum2 := numbers[i] + numbers[i-1] + numbers[i-2]
		if sum2 > sum1 {
			increased++
		}
	}
	return increased
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		s := scanner.Text()
		if i, err := strconv.Atoi(s); err == nil {
			data = append(data, i)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1")
	fmt.Println(Increased(data))
	fmt.Println("Part 2")
	fmt.Println(SlidingWindows(data))
}
