package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_input(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input_list []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		input_list = append(input_list, scanner.Text())
	}
	return input_list
}

func num_invalid(passwords []string) int {
	num_valid := 0
	for _, row := range passwords {
		policy_and_pass := strings.Split(row, ":")
		policy := strings.TrimSpace(policy_and_pass[0])
		password := strings.TrimSpace(policy_and_pass[1])
		char_limit := strings.Split(policy, " ")
		limited_char := char_limit[1]
		num_char_range := char_limit[0]
		min_and_max_num := strings.Split(num_char_range, "-")
		min, err := strconv.Atoi(min_and_max_num[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(min_and_max_num[1])
		if err != nil {
			log.Fatal(err)
		}

		limited_char_count := 0
		for _, char := range password {
			if string(char) == limited_char {
				limited_char_count++
			}
		}
		if limited_char_count >= min && limited_char_count <= max {
			num_valid++
		}
	}
	return num_valid
}

func main() {
	passwords := read_input("input.txt")
	fmt.Println(num_invalid(passwords))
}
