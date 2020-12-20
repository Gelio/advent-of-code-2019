package tile

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Rotation int

const (
	Hash = '#'
	Dot  = '.'
)

const (
	NoRotation Rotation = 0
	Rotate90   Rotation = 270
	Rotate180  Rotation = 180
	Rotate270  Rotation = 270
)

type Borders struct {
	Bottom string
	Top    string
	Right  string
	Left   string
}

type Tile struct {
	ID int
	Rotation
	Flipped bool
	Hashes  [][]bool
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

	t.Hashes = make([][]bool, len(lines)-1)
	for y, line := range lines[1:] {
		t.Hashes[y] = make([]bool, len(line))
		for x, c := range line {
			if c == Hash {
				t.Hashes[y][x] = true
			}
		}
	}

	t.fillBorders()

	return t, nil
}

func (t Tile) clone() Tile {
	newT := t

	newT.Hashes = make([][]bool, len(t.Hashes))
	for y, line := range t.Hashes {
		newT.Hashes[y] = make([]bool, len(line))
	}

	return newT
}

func (t Tile) flipVertical() Tile {
	newT := t.clone()
	newT.Flipped = true

	for y := range t.Hashes {
		for x := range t.Hashes[y] {
			newT.Hashes[len(t.Hashes)-1-y][x] = t.Hashes[y][x]
		}
	}
	newT.fillBorders()

	return newT
}

func (t Tile) rotate(rotation Rotation) Tile {
	newT := t.clone()
	sideLen := len(t.Hashes)
	newT.Rotation = rotation

	for y := range t.Hashes {
		for x := range t.Hashes[y] {
			if !t.Hashes[y][x] {
				continue
			}

			halfSide := float64(sideLen-1) / 2

			x1 := float64(x) - halfSide
			y1 := float64(y) - halfSide

			var angle = (float64(rotation) * math.Pi) / 180

			newX := x1*math.Cos(angle) - y1*math.Sin(angle)
			newY := x1*math.Sin(angle) + y1*math.Cos(angle)

			finalX := int(math.Round(newX + halfSide))
			finalY := int(math.Round(newY + halfSide))

			newT.Hashes[finalY][finalX] = true
		}
	}

	newT.fillBorders()

	return newT
}

func (t Tile) GetAllVariants() []Tile {
	variants := make([]Tile, 0, 8)

	variants = append(variants, t, t.rotate(Rotate90), t.rotate(Rotate180), t.rotate(Rotate270))

	flippedT := t.flipVertical()
	variants = append(variants, flippedT, flippedT.rotate(Rotate90), flippedT.rotate(Rotate180), flippedT.rotate(Rotate270))

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
