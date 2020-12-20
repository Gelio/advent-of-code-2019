package tile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	ID int
	Rotation
	Flip
	Borders Borders
}

func Parse(lines []string) (Tile, error) {
	idLineRegexp, err := regexp.Compile(`^Tile (\d+):$`)
	if err != nil {
		return Tile{}, fmt.Errorf("cannot compile tile ID line regexp: %w", err)
	}

	matches := idLineRegexp.FindStringSubmatch(lines[0])
	if len(matches) != 2 {
		return Tile{}, fmt.Errorf("first line (%q) does not match tile ID line", lines[0])
	}

	var t Tile
	t.ID, err = strconv.Atoi(matches[1])
	if err != nil {
		return t, fmt.Errorf("cannot parse tile ID (%q): %w", matches[1], err)
	}

	t.Borders = getBorders(lines[1:])

	return t, nil
}

func getBorders(img []string) Borders {
	var leftBorderSb, rightBorderSb strings.Builder
	var b Borders

	b.Top = img[0]
	b.Bottom = img[len(img)-1]

	for _, line := range img {
		leftBorderSb.WriteByte(line[0])
		rightBorderSb.WriteByte(line[len(line)-1])
	}

	b.Left = leftBorderSb.String()
	b.Right = rightBorderSb.String()

	return b
}
