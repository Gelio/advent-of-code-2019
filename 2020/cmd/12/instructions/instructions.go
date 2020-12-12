package instructions

import (
	"aoc-2020/cmd/12/ship"
	"fmt"
)

type Instruction interface {
	Exec(s *ship.Ship) error
}

type north struct {
	dist int
}

func (i north) Exec(s *ship.Ship) error {
	s.Y += i.dist
	return nil
}

type south struct {
	dist int
}

func (i south) Exec(s *ship.Ship) error {
	s.Y -= i.dist
	return nil
}

type east struct {
	dist int
}

func (i east) Exec(s *ship.Ship) error {
	s.X += i.dist
	return nil
}

type west struct {
	dist int
}

func (i west) Exec(s *ship.Ship) error {
	s.X -= i.dist
	return nil
}

type left struct {
	degrees int
}

func (i left) Exec(s *ship.Ship) error {
	s.Rotate(-i.degrees)
	return nil
}

type right struct {
	degrees int
}

func (i right) Exec(s *ship.Ship) error {
	s.Rotate(i.degrees)
	return nil
}

type forward struct {
	dist int
}

func (i forward) Exec(s *ship.Ship) error {
	if s.Rotation == 0 {
		north{i.dist}.Exec(s)
	} else if s.Rotation == 90 {
		east{i.dist}.Exec(s)
	} else if s.Rotation == 180 {
		south{i.dist}.Exec(s)
	} else if s.Rotation == 270 {
		west{i.dist}.Exec(s)
	} else {
		return fmt.Errorf("Invalid ship rotation when moving forward: %d", s.Rotation)
	}

	return nil
}
