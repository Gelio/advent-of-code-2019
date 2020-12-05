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
	reader := bufio.NewReader(os.Stdin)
	r, err := regexp.Compile("^(\\d+)-(\\d+) ([a-z]): (\\w+)$")

	if err != nil {
		fmt.Println("Cannot compile regexp", err)
		return
	}

	validPasswords := 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error while reading input", err)
			return
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		matchResult := r.FindStringSubmatch(line)

		if matchResult == nil {
			fmt.Println("Cannot match line:", line)
			return
		}

		position1, err := strconv.Atoi(matchResult[1])

		if err != nil {
			fmt.Println("Cannot parse position1", matchResult[1], "in line:", line)
			return
		}

		position2, err := strconv.Atoi(matchResult[2])

		if err != nil {
			fmt.Println("Cannot parse position2", matchResult[2], "in line:", line)
			return
		}

		characterToFind := matchResult[3][0]
		password := matchResult[4]

		if (password[position1-1] == characterToFind) != (password[position2-1] == characterToFind) {
			validPasswords++
		}
	}

	fmt.Println("Result: ", validPasswords)
}
