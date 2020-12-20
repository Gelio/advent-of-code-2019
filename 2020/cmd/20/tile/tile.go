package tile

import (
	"aoc-2020/cmd/20/positionhashmap"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	Hash = '#'
	Dot  = '.'
)

type Borders struct {
	Bottom string
	Top    string
	Right  string
	Left   string
}

type Tile struct {
	ID       int
	Rotation positionhashmap.Rotation
	Flipped  bool
	Hashes   positionhashmap.PositionHashMap
	Borders
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

	t.Hashes = positionhashmap.FromLines(lines[1:])

	t.fillBorders()

	return t, nil
}

func (t Tile) flipVertical() Tile {
	newT := t
	newT.Flipped = true
	newT.Hashes = t.Hashes.FlipVertical()
	newT.fillBorders()

	return newT
}

func (t Tile) rotate(rotation positionhashmap.Rotation) Tile {
	newT := t
	newT.Rotation = rotation
	newT.Hashes = t.Hashes.Rotate(rotation)

	newT.fillBorders()

	return newT
}

func (t Tile) GetAllVariants() []Tile {
	variants := make([]Tile, 0, 8)

	variants = append(variants, t, t.rotate(positionhashmap.Rotate90), t.rotate(positionhashmap.Rotate180), t.rotate(positionhashmap.Rotate270))

	flippedT := t.flipVertical()
	variants = append(variants, flippedT, flippedT.rotate(positionhashmap.Rotate90), flippedT.rotate(positionhashmap.Rotate180), flippedT.rotate(positionhashmap.Rotate270))

	return variants
}

func (t Tile) MatchesRight(otherT Tile) bool {
	return t.Borders.Right == otherT.Borders.Left
}

func (t Tile) MatchesBottom(otherT Tile) bool {
	return t.Borders.Bottom == otherT.Borders.Top
}

func (t *Tile) fillBorders() {
	var topBorderSb, bottomBorderSb, leftBorderSb, rightBorderSb strings.Builder

	for x := range t.Hashes[0] {
		topBorderSb.WriteRune(hashToRune(t.Hashes[0][x]))
		bottomBorderSb.WriteRune(hashToRune(t.Hashes[len(t.Hashes)-1][x]))
	}
	t.Borders.Top = topBorderSb.String()
	t.Borders.Bottom = bottomBorderSb.String()

	for _, line := range t.Hashes {
		leftBorderSb.WriteRune(hashToRune(line[0]))
		rightBorderSb.WriteRune(hashToRune(line[len(line)-1]))
	}

	t.Borders.Left = leftBorderSb.String()
	t.Borders.Right = rightBorderSb.String()
}

func hashToRune(v bool) rune {
	if v {
		return Hash
	}

	return Dot
}
