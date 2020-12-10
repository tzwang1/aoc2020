package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_inputs(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input_list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_list = append(input_list, scanner.Text())
	}
	return input_list
}

func count_valid_passwords(passwords []string) int {
	num_valid := 0
	for _, row := range passwords {
		policy_and_password := strings.Split(row, ":")
		policy := strings.TrimSpace(policy_and_password[0])
		password := strings.TrimSpace(policy_and_password[1])
		char_limit := strings.Split(policy, " ")
		limited_char := char_limit[1]
		start_and_end := strings.Split(char_limit[0], "-")
		start, err := strconv.Atoi(start_and_end[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(start_and_end[1])
		if err != nil {
			log.Fatal(err)
		}

		start_char := string(password[start-1])
		end_char := string(password[end-1])

		if (start_char == limited_char || end_char == limited_char) && !(start_char == limited_char && end_char == limited_char) {
			num_valid++
		}
	}
	return num_valid
}

func main() {
	passwords := read_inputs("input.txt")
	fmt.Println(count_valid_passwords(passwords))
}
