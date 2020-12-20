package stdin

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadAllLines reads all lines from stdin
func ReadAllLines() ([]string, error) {
	return ReadLinesFromReader(os.Stdin)
}

// ReadLinesFromFile reads all lines from an input file
func ReadLinesFromFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", name, err)
	}

	return ReadLinesFromReader(f)
}

// ReadLinesFromReader reads lines from a given reader
func ReadLinesFromReader(ioR io.Reader) ([]string, error) {
	var lines []string
	r := bufio.NewReader(ioR)

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					lines = append(lines, strings.TrimSpace(line))
				}
				break
			}

			return nil, err
		}

		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	return lines, nil
}
