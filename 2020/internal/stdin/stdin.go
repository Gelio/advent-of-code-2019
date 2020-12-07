package stdin

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadAllLines reads all lines from stdin
func ReadAllLines() (lines []string, err error) {
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
