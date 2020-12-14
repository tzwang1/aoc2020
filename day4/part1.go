package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "strconv"
)

func split_password_fields(passports []string) (password_maps []map[string]string) {
	for _, passport := range passports {
		passport = strings.TrimSpace(passport)
		fields := strings.Split(passport, " ")
		stored_fields := make(map[string]string)

		for _, field := range fields {
			split_field := strings.Split(field, ":")
			stored_fields[split_field[0]] = split_field[1]
		}
		password_maps = append(password_maps, stored_fields)
	}
	return
}

func num_valid_passports(passports_map []map[string]string, required_fields []string) int {
	num_valid := 0
	for _, passport_map := range passports_map {
		is_valid := true
		for _, required_field := range required_fields {
			if _, ok := passport_map[required_field]; !ok {
				is_valid = false
				break
			}
		}

		if is_valid && fields_valid(passport_map) {
			num_valid++
		}
	}
	return num_valid
}

func birth_year_valid(byr string) bool {
    ibyr, err := strconv.Atoi(byr)
    if err != nil {
        return false
    }
    return ibyr >= 1920 && ibyr <= 2002
}

func issue_year_valid(iyr string) bool {
    iiyr, err := strconv.Atoi(iyr)
    if err != nil {
        return false
    }
    return iiyr >= 2010 && iiyr <= 2020
}

func expiration_year_valid(eyr string) bool {
    ieyr, err := strconv.Atoi(eyr)
    if err != nil { 
        return false
    }
    return ieyr >= 2020 && ieyr <= 2030
}

func height_valid(height string) bool {
    if strings.HasSuffix(height, "cm") { 
        height_val := height[:len(height)-2]
        height_cm, err := strconv.Atoi(height_val)
        if err != nil {
            return false
        }
        return height_cm >= 150 && height_cm <= 193
    } else if strings.HasSuffix(height, "in") {
        height_val := height[:len(height)-2]
        height_in, err := strconv.Atoi(height_val)
        if err != nil {
            return false
        }
        return height_in >= 59 && height_in <= 76
    } else {
        return false
    }
}

func haircolor_valid(haircolor string) bool {
    if string(haircolor[0]) != "#" {
        return false
    }
    for i := 1; i < len(haircolor); i++ {
        if (haircolor[i] < 48 || haircolor[i] > 57) && (haircolor[i] < 97 || haircolor[i] > 102) {
            return false
        }
    }
    return true
}

func eyecolor_valid(eyecolor string) bool {
    valid_eyecolor := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true} 

    if _, ok := valid_eyecolor[eyecolor]; ok {
        return true
    }
    return false
}

func passport_id_valid(passport_id string) bool {
    if len(passport_id) != 9 {
        return false
    }
    _, err := strconv.Atoi(passport_id)
    if err != nil {
        return false
    }
    return true
}
    

func fields_valid(passport_map map[string]string) bool {
    return birth_year_valid(passport_map["byr"]) &&
            issue_year_valid(passport_map["iyr"]) &&
            expiration_year_valid(passport_map["eyr"]) &&
            height_valid(passport_map["hgt"]) &&
            haircolor_valid(passport_map["hcl"]) &&
            eyecolor_valid(passport_map["ecl"]) &&
            passport_id_valid(passport_map["pid"])
}

func read_input(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input_line []string
	scanner := bufio.NewScanner(file)
	cur_line := ""
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			cur_line = cur_line + scanner.Text() + " "
		} else {
			input_line = append(input_line, cur_line)
			cur_line = ""
		}
	}
    if cur_line != "" {
        input_line = append(input_line, cur_line)
    }
	return input_line
}

func main() {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	passports := read_input("input.txt")
	passport_maps := split_password_fields(passports)
	fmt.Println(num_valid_passports(passport_maps, required_fields))

}
