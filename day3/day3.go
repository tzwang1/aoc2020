package main

import (
    "bufio"
    "fmt"
    "os"
)

type movement struct {
    x int
    y int
}

func read_input(input string) [][]string {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var input_matrix [][]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var line []string
        for _, c := range scanner.Text() {
            line = append(line, string(c))
        }
        input_matrix = append(input_matrix, line)
    }
    return input_matrix
}

func count_num_trees(matrix [][]string, right int, down int) int {
    num_trees := 0
    cur_x := 0
    cur_y := 0

    for cur_x < len(matrix) {
        cur_x = cur_x + down
        cur_y = cur_y + right
        cur_y = cur_y % len(matrix[0])
        if cur_x < len(matrix) &&  matrix[cur_x][cur_y] == "#" {
            num_trees++
        }
    }
    return num_trees
}

func main() {
    input_matrix := read_input("input.txt")
    fmt.Println(count_num_trees(input_matrix, 3, 1) * count_num_trees(input_matrix, 1, 1) * count_num_trees(input_matrix, 5, 1) * count_num_trees(input_matrix, 7, 1) * count_num_trees(input_matrix, 1, 2))
}
