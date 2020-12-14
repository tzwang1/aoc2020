package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

type boarding_pass struct {
    row string
    col string
}

func calculate_row(row_code string, min int, max int) int {
    for _, code := range row_code {
        mid_point := min + ((max - min) / 2)
       if string(code) == "F" {
            max = mid_point
       } else {
            min = mid_point
       }
    }
    return max
}

func calculate_col(col_code string, min int, max int) int {
    for _, code := range col_code {
        mid_point := min + ((max - min) / 2)
        if string(code) == "L" {
            max = mid_point
        } else {
            min = mid_point
        }
    }
    return max
}

func calculate_seat_id(col int, row int) int {
   return row * 8 + col 
}

func part1(boarding_passes []boarding_pass) int {
    max_seat_id := 0
    for _, pass := range boarding_passes {
        row := calculate_row(pass.row, 0, 127)
        col := calculate_col(pass.col, 0, 7)
        seat_id := calculate_seat_id(col, row)
        if seat_id > max_seat_id {
            max_seat_id = seat_id
        }
    }
    return max_seat_id
}

func get_seat_ids(boarding_passes []boarding_pass) []int {
    var seat_ids []int 
    for _, pass := range boarding_passes {
        row := calculate_row(pass.row, 0, 127)
        col := calculate_col(pass.col, 0, 7)
        seat_id := calculate_seat_id(col, row)
        seat_ids = append(seat_ids, seat_id)
    }
    return seat_ids
}

func part2(seat_ids []int) int {
    sort.Ints(seat_ids)
    for i := 0; i < len(seat_ids)-1; i++ {   
        if seat_ids[i] == seat_ids[i+1]-2 {
            return seat_ids[i]+1
        }
    }
    return -1
}

func read_input(input string) []boarding_pass {
    file, err := os.Open(input)

    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input_lines []boarding_pass

    for scanner.Scan() {
        line := scanner.Text()
        input_lines = append(input_lines, boarding_pass{line[:7], line[len(line)-3:]})
    }
    return input_lines
}

func main() {
    boarding_passes := read_input("input.txt")
    fmt.Println(part1(boarding_passes))
    seat_ids := get_read_ids(boarding_passes)
    fmt.Println(part2(seat_ids))
}
