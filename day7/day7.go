package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type BagEdge struct {
    outer_bag string
    inner_bag string
}

func read_input(input string) map[string][]string {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    var bag_edges []BagEdge
    for scanner.Scan() {
        line := scanner.Text()
        
        outer_and_inner := strings.Split(line, "contain")
        outer_bag := outer_and_inner[0]
        outer_bag = strings.TrimSpace(outer_bag)
        outer_bag = outer_bag[:len(outer_bag)-1]

        inner_bags := strings.Split(outer_and_inner[1], ",")
        for _, inner_bag := range inner_bags {
            inner_bag_clean := strings.Trim(strings.TrimSpace(inner_bag[2:]), ".")
            if strings.HasSuffix(inner_bag_clean, "bags") {
                inner_bag_clean = inner_bag_clean[:len(inner_bag_clean)-1]
            }
            bag_edges = append(bag_edges, BagEdge{outer_bag, inner_bag_clean})
        }
    }

    bag_relationship := make(map[string][]string)
    for _, bag_edge := range bag_edges {
        bag_relationship[bag_edge.inner_bag] =
        append(bag_relationship[bag_edge.inner_bag], bag_edge.outer_bag)
    }
    return bag_relationship
}

func find_num_outer_bags(bag_relationship map[string][]string) int {
    num_bags := 0
    seen_bags := make(map[string]bool)
    stack := []string{"shiny gold bag"}
    
    for len(stack) != 0 {
        cur_bag := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        for _, inner_bag := range bag_relationship[cur_bag] {
            if _, ok := seen_bags[inner_bag]; ok {
                continue
            }
            seen_bags[inner_bag] = true
            num_bags++
            stack = append(stack, inner_bag)
        }
    }
    return num_bags
}

func main() {
    bag_relationship := read_input("input.txt")
    fmt.Println(bag_relationship)
    fmt.Println(find_num_outer_bags(bag_relationship))
}
