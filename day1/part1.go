package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func read_input(input string) []int {
    file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    var input_list []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        input_list = append(input_list, num)
    }
    return input_list
}

func solve(entry_list []int) int {
    seen_entries := make(map[int]bool)

    for _, entry := range entry_list {
        if _, ok := seen_entries[2020 - entry]; ok {
            return entry * (2020 - entry)
        }
        seen_entries[entry] = true
    }
    return -1
}

func main() {
    entry_list := read_input("input1.txt")
    fmt.Println(solve(entry_list))
}
