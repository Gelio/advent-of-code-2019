package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, err := readAllLines()

	if err != nil {
		fmt.Println("Error while reading input", err)
		return
	}

	passportLines := getPassportLines(lines)

	validPassports := 0

	for _, singlePassportLines := range passportLines {
		passportFields, err := parsePassportFields(singlePassportLines)

		if err != nil {
			fmt.Println("Error parsing passport fields", err)
			return
		}

		if isPassportValid(passportFields) {
			validPassports++
		}
	}

	fmt.Println("Result: ", validPassports)
}

func readAllLines() (lines []string, err error) {
	reader := bufio.NewReader(os.Stdin)
	var line string

	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}

			return
		}

		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	return
}

func getPassportLines(lines []string) (passports [][]string) {
	var currentPassport []string

	for _, line := range lines {
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = make([]string, 0)
			continue
		}

		currentPassport = append(currentPassport, line)
	}

	if len(currentPassport) > 0 {
		passports = append(passports, currentPassport)
	}

	return
}

func parsePassportFields(passportLines []string) (passportFields map[string]string, err error) {
	passportFields = make(map[string]string)
	for _, line := range passportLines {
		fields := strings.Split(line, " ")

		for _, field := range fields {
			fieldParts := strings.Split(field, ":")
			if len(fieldParts) != 2 {
				return nil, fmt.Errorf("Invalid field %s (does not contain 2 parts)", field)
			}

			fieldName := fieldParts[0]
			fieldValue := fieldParts[1]
			passportFields[fieldName] = fieldValue
		}
	}

	return
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var validationRules = map[string]func(string) bool{
	"byr": func(value string) bool {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return num >= 1920 && num <= 2002
	},
	"iyr": func(value string) bool {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return num >= 2010 && num <= 2020
	},
	"eyr": func(value string) bool {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return num >= 2020 && num <= 2030
	},
	"hgt": func(value string) bool {
		if len(value) < 2 {
			return false
		}
		unit := value[len(value)-2:]
		height := value[:len(value)-2]
		num, err := strconv.Atoi(height)
		if err != nil {
			return false
		}

		switch unit {
		case "cm":
			return num >= 150 && num <= 193
		case "in":
			return num >= 59 && num <= 76
		default:
			return false
		}
	},
	"hcl": func(value string) bool {
		r, err := regexp.Compile("^#[0-9a-f]{6}$")
		if err != nil {
			fmt.Println("Error when compiling regexp")
			panic(err)
		}

		return r.MatchString(value)
	},
	"ecl": func(value string) bool {
		allowedColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

		return allowedColors[value]
	},
	"pid": func(value string) bool {
		r, err := regexp.Compile("^\\d{9}$")
		if err != nil {
			fmt.Println("Error when compiling regexp")
			panic(err)
		}

		return r.MatchString(value)
	},
}

func isPassportValid(passportFields map[string]string) bool {
	for fieldName, isFieldValid := range validationRules {
		fieldValue, ok := passportFields[fieldName]
		if !ok {
			return false
		}

		if !isFieldValid(fieldValue) {
			return false
		}
	}

	return true
}
