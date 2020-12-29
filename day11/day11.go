package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(input string) [][]string {
	file, err := os.Open(input)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seat_arrangement := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		cur_row := []string{}

		for _, char := range line {
			cur_row = append(cur_row, string(char))
		}
		seat_arrangement = append(seat_arrangement, cur_row)
	}
	return seat_arrangement
}

func iterate(seating_arrangement [][]string) int {
	stabilized := false
	for !stabilized {
		temp_arrangement := [][]string{}
		for i := 0; i < len(seating_arrangement); i++ {
			temp_row := []string{}
			for j := 0; j < len(seating_arrangement[i]); j++ {
				cur_seat := seating_arrangement[i][j]
				if cur_seat == "L" {
					no_occupied := true
					for _, neighbour := range getNeighbours(i, j) {
						n_row := neighbour[0]
						n_col := neighbour[1]
						if isValid(n_row, n_col, seating_arrangement) {
							if seating_arrangement[n_row][n_col] == "#" {
								no_occupied = false
								break
							}
						}
					}
					if no_occupied {
						temp_row = append(temp_row, "#")
					} else {
						temp_row = append(temp_row, cur_seat)
					}
				} else if cur_seat == "#" {
					num_occupied := 0
					for _, neighbour := range getNeighbours(i, j) {
						n_row := neighbour[0]
						n_col := neighbour[1]
						if isValid(n_row, n_col, seating_arrangement) {
							if seating_arrangement[n_row][n_col] == "#" {
								num_occupied++
							}
						}
					}
					if num_occupied >= 4 {
						temp_row = append(temp_row, "L")
					} else {
						temp_row = append(temp_row, cur_seat)
					}
				} else {
					temp_row = append(temp_row, cur_seat)
				}
			}
			temp_arrangement = append(temp_arrangement, temp_row)
		}
		equal := true
		for i := 0; i < len(temp_arrangement); i++ {
			for j := 0; j < len(temp_arrangement[i]); j++ {
				if seating_arrangement[i][j] != temp_arrangement[i][j] {
					equal = false
				}
				seating_arrangement[i][j] = temp_arrangement[i][j]
			}
		}

		if equal {
			stabilized = true
		}
	}

	num_occupied := 0
	for i := 0; i < len(seating_arrangement); i++ {
		for j := 0; j < len(seating_arrangement[i]); j++ {
			if seating_arrangement[i][j] == "#" {
				num_occupied++
			}
		}
	}
	return num_occupied
}

func isValid(row, col int, matrix [][]string) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[row])
}

func getNeighbours(row, col int) [][]int {
	neighbours := [][]int{}
	neighbours = append(neighbours, []int{row - 1, col - 1})
	neighbours = append(neighbours, []int{row, col - 1})
	neighbours = append(neighbours, []int{row + 1, col - 1})
	neighbours = append(neighbours, []int{row + 1, col})
	neighbours = append(neighbours, []int{row + 1, col + 1})
	neighbours = append(neighbours, []int{row, col + 1})
	neighbours = append(neighbours, []int{row - 1, col + 1})
	neighbours = append(neighbours, []int{row - 1, col})

	return neighbours
}

func getOccupiedNeighbours(row, col int, seating_arrangement [][]string) [][]int {
	neighbours := [][]int{}
	neighbours = append(neighbours, checkDirection(row, col, -1, -1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, 0, -1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, 1, -1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, 1, 0, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, 1, 1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, 0, 1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, -1, 1, seating_arrangement))
	neighbours = append(neighbours, checkDirection(row, col, -1, 0, seating_arrangement))

	return neighbours
}

func checkDirection(start_row, start_col, row_move, col_move int, seating_arrangement [][]string) []int {
	cur_pos := []int{start_row + row_move, start_col + col_move}
	for isValid(cur_pos[0], cur_pos[1], seating_arrangement) &&
		seating_arrangement[cur_pos[0]][cur_pos[1]] == "." {
		cur_pos[0] += row_move
		cur_pos[1] += col_move
	}
	return cur_pos
}

func iteratePart2(seating_arrangement [][]string) int {
	stabilized := false
	for !stabilized {
		temp_arrangement := [][]string{}
		for i := 0; i < len(seating_arrangement); i++ {
			temp_row := []string{}
			for j := 0; j < len(seating_arrangement[i]); j++ {
				cur_seat := seating_arrangement[i][j]
				if cur_seat == "L" {
					no_occupied := true
					for _, neighbour := range getOccupiedNeighbours(i, j, seating_arrangement) {
						n_row := neighbour[0]
						n_col := neighbour[1]
						if isValid(n_row, n_col, seating_arrangement) {
							if seating_arrangement[n_row][n_col] == "#" {
								no_occupied = false
								break
							}
						}
					}
					if no_occupied {
						temp_row = append(temp_row, "#")
					} else {
						temp_row = append(temp_row, cur_seat)
					}
				} else if cur_seat == "#" {
					num_occupied := 0
					for _, neighbour := range getOccupiedNeighbours(i, j, seating_arrangement) {
						n_row := neighbour[0]
						n_col := neighbour[1]
						if isValid(n_row, n_col, seating_arrangement) {
							if seating_arrangement[n_row][n_col] == "#" {
								num_occupied++
							}
						}
					}
					if num_occupied >= 5 {
						temp_row = append(temp_row, "L")
					} else {
						temp_row = append(temp_row, cur_seat)
					}
				} else {
					temp_row = append(temp_row, cur_seat)
				}
			}
			temp_arrangement = append(temp_arrangement, temp_row)
		}
		equal := true
		for i := 0; i < len(temp_arrangement); i++ {
			for j := 0; j < len(temp_arrangement[i]); j++ {
				if seating_arrangement[i][j] != temp_arrangement[i][j] {
					equal = false
				}
				seating_arrangement[i][j] = temp_arrangement[i][j]
			}
		}

		if equal {
			stabilized = true
		}
	}

	num_occupied := 0
	for i := 0; i < len(seating_arrangement); i++ {
		for j := 0; j < len(seating_arrangement[i]); j++ {
			if seating_arrangement[i][j] == "#" {
				num_occupied++
			}
		}
	}
	return num_occupied
}

func main() {
	seating_arrangement := readInput("input.txt")
//	fmt.Println(iterate(seating_arrangement))
    fmt.Println(iteratePart2(seating_arrangement))
}
