package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func readInput(input string) (adapters []int) {
    file, err := os.Open(input)

    if err != nil {
        panic(err)
    }
    defer file.Close()
    adapters = append(adapters, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        
        if err != nil {
            panic(err)
        }
        adapters = append(adapters, num)
    }
    sort.Ints(adapters)
    adapters = append(adapters, adapters[len(adapters)-1]+3)
    return adapters
}

func findDifferences(adapters []int) int {
    adapter_differences := map[int]int{}

    for i := 0; i < len(adapters)-1; i++ {
        cur_diff := adapters[i+1] - adapters[i]
        if _, ok := adapter_differences[cur_diff]; ok {
            adapter_differences[cur_diff]++
        } else {
            adapter_differences[cur_diff] = 1
        }
    }
    fmt.Println(adapter_differences)
    return adapter_differences[1] * adapter_differences[3]
}

func buildGraph(adapters[]int) map[int][]int {
    graph := map[int][]int{}

    for i := 0; i < len(adapters); i++ {
        for j := i+1; j < len(adapters); j++ {
            if adapters[j] - adapters[i] <= 3 {
               graph[adapters[i]] = append(graph[adapters[i]], adapters[j]) 
            }
        }
    }
    return graph
}

func findNumDistinctAdapterPathsIterative(graph map[int][]int, target int) int {
    num_ways = []int{1}

    for i := target - 1; i >=0 ; i-- {
        // Look at previous
    }
}

func findNumDistinctAdapterPathsRecursive(graph map[int][]int, target int) int {
    cur_path := []int{0}
    seen := map[int]int{}
    return helper(graph, target, &cur_path, &seen)
}

func helper(graph map[int][]int, target int, cur_path *[]int, seen *map[int]int) int {
    cur_adapter := (*cur_path)[len(*cur_path)-1]
    num_paths := 0
    if cur_adapter== target {
        return 1
    } else if _, ok := (*seen)[cur_adapter]; ok {
        return (*seen)[cur_adapter]
    } else {
        for _, next_adapter := range graph[cur_adapter] {
            *cur_path = append(*cur_path, next_adapter)
            num_paths += helper(graph, target, cur_path, seen)
            *cur_path = (*cur_path)[:len(*cur_path)-1]
        }
    }
    (*seen)[cur_adapter] = num_paths
    return num_paths
}

func main() {
    adapters := readInput("input.txt")
    fmt.Println(adapters)
    target := adapters[len(adapters)-1]
    fmt.Println(findDifferences(adapters))
    graph := buildGraph(adapters)
    fmt.Println(graph)
    fmt.Println(findNumDistinctAdapterPathsRecursive(graph, target))
}
