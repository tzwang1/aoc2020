package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Instruction struct {
    operation string
    num int
}

func readInput(input string) (instructions []Instruction) {
    file, err := os.Open(input)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        operation_and_num := strings.Split(line, " ")
        operation := operation_and_num[0]
        num_str := operation_and_num[1]
        num, err := strconv.Atoi(num_str[1:])
        if err != nil {
            panic(err)
        }
        if num_str[0] == '-' {
            num = -1 * num
        }
        instructions = append(instructions, Instruction{operation, num}) 
    }
    return 
}

func fixInstructions(instructions []Instruction) []Instruction {
    can_reach_last_instruction := make(map[int]bool)
    can_reach_last_instruction[len(instructions)] = true
    none_can_reach := false
    for !none_can_reach {
        none_can_reach = true
        for idx, instruction := range instructions {
            if instruction.operation == "jmp" {
                if _, ok := can_reach_last_instruction[idx+instruction.num]; ok {
                    if _, ok := can_reach_last_instruction[idx]; !ok {
                        can_reach_last_instruction[idx] = true
                        none_can_reach = false
                    }
                }
            } else {
                if _, ok := can_reach_last_instruction[idx+1]; ok {
                    if _, ok := can_reach_last_instruction[idx]; !ok {
                        can_reach_last_instruction[idx] = true
                        none_can_reach = false
                    }
                }
            }
        }
        fmt.Println(can_reach_last_instruction)
        fmt.Println("")
    }
    fmt.Println(can_reach_last_instruction)
    
    idx := 0
    for idx < len(instructions) {
        idx_str := strconv.Itoa(idx)
        instruction := instructions[idx]
        if instruction.operation == "acc" {
            idx++
            continue
        } else if instruction.operation == "nop" {
            if _, ok := can_reach_last_instruction[idx + instruction.num]; ok {
                instructions[idx].operation = "jmp"
                fmt.Println("Setting instruction at idx: " + idx_str + " to jmp")
                break
            }
            idx++
        } else {
            if _, ok := can_reach_last_instruction[idx+1]; ok {
                instructions[idx].operation = "nop"
                fmt.Println("Setting instruction at idx: " + idx_str + " to nop")
                break
            }
            idx+=instruction.num

        }
    }
    return instructions
}

func findAccValueBeforeInfiniteLoop(instructions []Instruction) (acc_val int) {
    cur_idx := 0
    seen_instruction_idx := make(map[int]bool)
    seen := false

    for !seen && cur_idx < len(instructions) {
        if _, ok := seen_instruction_idx[cur_idx]; ok {
            seen = true
            break
        }
        seen_instruction_idx[cur_idx] = true
        instruction := instructions[cur_idx]
        if instruction.operation == "nop" {
            cur_idx++
        } else if instruction.operation == "acc" {
            acc_val += instruction.num
            cur_idx++
        } else {
            cur_idx += instruction.num
        }
    }
    return
}

func main() {
    instructions := readInput("input.txt")
    instructions = fixInstructions(instructions)
    for idx, instruction := range instructions {
        if idx == 47 {
            fmt.Println("Index: " + strconv.Itoa(idx), instruction)
            fmt.Println("")
        }
    }
    fmt.Println(findAccValueBeforeInfiniteLoop(instructions))
}
