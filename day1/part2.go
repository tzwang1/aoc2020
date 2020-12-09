package main 

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
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
    sort.Ints(entry_list)

    for idx, entry := range entry_list {
        start := 0
        end := len(entry_list)-1

        for start < end {
            if start == idx {
                start++
                continue
            }
            if end == idx {
                end--
                continue
            }
            cur_sum := entry + entry_list[end] + entry_list[start]
            if cur_sum == 2020 {
                return entry * entry_list[end] * entry_list[start]
            } else if cur_sum < 2020 {
                start++
            } else {
                end--
            }
        }
    }
    return -1
}

func main() {
    entry_list := read_input("input1.txt")
    fmt.Println(solve(entry_list))
}
