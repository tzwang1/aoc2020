package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

type BagRelationship struct {
    outer string
    inner string
    num_contains int
}

type Bag struct {
    bag_type string
    num int
}

func read_input(input string) map[string][]Bag {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)

    var bag_relationship []BagRelationship
    for scanner.Scan() {
        line := scanner.Text()

        outer_and_inner := strings.Split(line, "contain")
        outer_bag := outer_and_inner[0]
        outer_bag = strings.TrimSpace(outer_bag)
        outer_bag = outer_bag[:len(outer_bag)-1]

        inner_bags := strings.Split(outer_and_inner[1], ",")
        for _, inner_bag := range inner_bags {
            fmt.Println(inner_bag)
            inner_bag_clean := strings.Trim(strings.TrimSpace(inner_bag[2:]), ".")
            if strings.HasSuffix(inner_bag_clean, "bags") {
                inner_bag_clean = inner_bag_clean[:len(inner_bag_clean)-1]
            }
            num_bag, err := strconv.Atoi(string(strings.TrimSpace(inner_bag)[0]))
            if err != nil {
                continue
            }

            bag_relationship = append(bag_relationship, BagRelationship{outer_bag, inner_bag_clean, num_bag})
        }
    }
    bag_graph := make(map[string][]Bag)
    for _, relationship := range bag_relationship {
        bag_graph[relationship.outer] = append(bag_graph[relationship.outer], Bag{relationship.inner, relationship.num_contains})
    }
    return bag_graph
}

func count_num_bags(bag_graph map[string][]Bag, bag string) int {
    count := 0
    stack := []Bag{Bag{bag, 1}}
   // seen_bags := make(map[string]bool)

    for len(stack) != 0 {
        cur_bag := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        bag_type := cur_bag.bag_type
        count_outer := cur_bag.num
        fmt.Println(bag_type)
        fmt.Println("Counter outer: ", count_outer)
        fmt.Println("")
        for _, inner_bag := range bag_graph[bag_type] {
     //       if _, ok := seen_bags[inner_bag.bag_type]; ok {
     //           continue
     //       }
    //        seen_bags[inner_bag.bag_type] = true
            count += count_outer * inner_bag.num
            fmt.Println("\tInner bag: ", inner_bag.bag_type)
            fmt.Println("\tInner bag count: ", inner_bag.num)
            fmt.Println("\tTotal count: ", count)
            stack = append(stack, Bag{inner_bag.bag_type, count_outer * inner_bag.num})
        }
    }
    return count
}

func main() {
    bag_graph := read_input("input.txt")
    for key, value := range bag_graph {
        fmt.Println(key, " contains ", value)
    }
    fmt.Println(count_num_bags(bag_graph, "shiny gold bag"))
}
