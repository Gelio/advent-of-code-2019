package positionhashmap

import "math"

type Rotation int

const (
	NoRotation Rotation = 0
	Rotate90   Rotation = 90
	Rotate180  Rotation = 180
	Rotate270  Rotation = 270
)

type PositionHashMap [][]bool

func FromLines(lines []string) PositionHashMap {
	phm := make(PositionHashMap, len(lines))

	for y, line := range lines {
		phm[y] = make([]bool, len(line))

		for x, c := range line {
			if c == '#' {
				phm[y][x] = true
			}
		}
	}

	return phm
}

func New(sideLen int) PositionHashMap {
	phm := make(PositionHashMap, sideLen)
	for y := 0; y < sideLen; y++ {
		phm[y] = make([]bool, sideLen)
	}

	return phm
}

func (phm PositionHashMap) FlipVertical() PositionHashMap {
	newPhm := New(len(phm))

	for y := range phm {
		for x := range phm[y] {
			newPhm[len(phm)-1-y][x] = phm[y][x]
		}
	}

	return newPhm
}

func (phm PositionHashMap) Rotate(rotation Rotation) PositionHashMap {
	sideLen := len(phm)
	newPhm := New(sideLen)

	for y := range phm {
		for x := range phm[y] {
			if !phm[y][x] {
				continue
			}

			halfSide := float64(sideLen-1) / 2

			x1 := float64(x) - halfSide
			y1 := float64(y) - halfSide

			angle := (float64(rotation) * math.Pi) / 180

			newX := x1*math.Cos(angle) - y1*math.Sin(angle)
			newY := x1*math.Sin(angle) + y1*math.Cos(angle)

			finalX := int(math.Round(newX + halfSide))
			finalY := int(math.Round(newY + halfSide))

			newPhm[finalY][finalX] = true
		}
	}

	return newPhm
}

func (phm PositionHashMap) GetAllVariants() []PositionHashMap {
	variants := make([]PositionHashMap, 0, 8)

	variants = append(variants, phm, phm.Rotate(Rotate90), phm.Rotate(Rotate180), phm.Rotate(Rotate270))

	flippedPhm := phm.FlipVertical()
	variants = append(variants, flippedPhm, flippedPhm.Rotate(Rotate90), flippedPhm.Rotate(Rotate180), flippedPhm.Rotate(Rotate270))

	return variants
}

func (phm PositionHashMap) Contains(otherPhm PositionHashMap, offsetX, offsetY int) bool {
	for y, row := range otherPhm {
		for x, c := range row {
			if !c {
				continue
			}

			if !phm[offsetY+y][offsetX+x] {
				return false
			}
		}
	}

	return true
}

func (phm PositionHashMap) CountValues() int {
	sum := 0
	for _, row := range phm {
		for _, c := range row {
			if c {
				sum++
			}
		}
	}

	return sum
}
