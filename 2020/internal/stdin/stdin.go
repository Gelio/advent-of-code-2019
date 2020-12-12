package stdin

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadAllLines reads all lines from stdin
func ReadAllLines() ([]string, error) {
	r := bufio.NewReader(os.Stdin)

	return readLinesFromReader(r)
}

// ReadLinesFromFile reads all lines from an input file
func ReadLinesFromFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)

	return readLinesFromReader(r)
}

func readLinesFromReader(r *bufio.Reader) ([]string, error) {
	var lines []string

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	return lines, nil
}
