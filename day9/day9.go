package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func readInput(input string) (nums []int) {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cur_num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            panic(err)
        }
        nums = append(nums, cur_num)
    }
    return nums
}

func findWeakness(nums []int, preamble_size int) int {
    cur_preamble_map := make(map[int]bool)
    for i := 0; i < preamble_size; i++ {
        cur_preamble_map[nums[i]] = true
    }
    cur_idx := preamble_size
    fmt.Println(cur_preamble_map)

    for cur_idx < len(nums) {
        cur_num := nums[cur_idx]
        found := false
        for i := cur_idx - preamble_size; i < cur_idx; i++ {
            fmt.Println(cur_num - nums[i])
            if _, ok := cur_preamble_map[cur_num - nums[i]]; ok {
                found = true
                break
            }
        }
        if !found {
            return cur_num
        }
        cur_preamble_map[cur_num] = true
        delete(cur_preamble_map, nums[cur_idx-preamble_size])
        cur_idx++
    }
    return -1
}

func findEncryptionWeakness(weak_num int, nums []int) int {
    for i := 0; i < len(nums); i++ {
        cur_sum := nums[i]
        for j := i+1; j < len(nums); j++ {
            cur_sum+=nums[j]
            if cur_sum == weak_num {
                weak_range := nums[i:j+1]
                sort.Ints(weak_range)
                return weak_range[0] + weak_range[len(weak_range)-1]
            }
        }
    }
    return -1
}

func main() {
    nums := readInput("input.txt")
    weak_num := findWeakness(nums, 25)
    fmt.Println(findEncryptionWeakness(weak_num, nums))
}
