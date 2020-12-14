package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func read_input(input string) []string {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var input_lines []string
    scanner := bufio.NewScanner(file)

    cur_group := ""
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            cur_group = cur_group + line
            fmt.Println("Cur group: ", cur_group)
        } else {
            input_lines = append(input_lines, cur_group)
            cur_group = ""
        }
    }
    if cur_group != "" {
        input_lines = append(input_lines, cur_group)
    }
    return input_lines
}

func read_input2(input string) []string {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var input_lines []string
    scanner := bufio.NewScanner(file)

    cur_group := ""
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            cur_group = cur_group + line + ","
        } else {
            input_lines = append(input_lines, cur_group)
            cur_group = ""
        }
    }
    if cur_group != "" {
        input_lines = append(input_lines, cur_group)
    }
    return input_lines
}

func part1(groups []string) int {
    total_yes := 0
    for _, group := range groups {
        seen := make(map[string]bool)
        for _, ans := range group {
            seen[string(ans)] = true
        }
        total_yes += len(seen)
    }
    return total_yes
}

func part2(groups []string) int {
    total_all_yes := 0
    for _, group := range groups {
        group = strings.Trim(group, ",")
        ans := strings.Split(group, ",") 
        fmt.Println(group)
        for _, char := range ans[0] {
            char_in := true
            for i := 0; i < len(ans); i++ {
               if !strings.Contains(ans[i], string(char)) {
                    char_in = false
               }
            }
            if char_in {
                total_all_yes++
            } else {
                fmt.Println("Cur Ans: ", ans)
                fmt.Println("Group: ", group)
            }
        }
    }
    return total_all_yes
}

func main() {
    groups := read_input("input.txt")
    fmt.Println(part1(groups))
    groups2 := read_input2("input.txt")
    fmt.Println(part2(groups2))
}
