package instructions

import (
	"aoc-2020/cmd/12/ship"
	"aoc-2020/cmd/12/waypoint"
	"fmt"
)

type Instruction interface {
	// ExecShip executes the instruction on the ship (from part 1)
	ExecShip(s *ship.Ship) error
	// ExecWaypoint executes the instruction based on waypoint rules (from part 2)
	ExecWaypoint(w *waypoint.Waypoint) error
}

type north struct {
	dist int
}

func (i north) ExecShip(s *ship.Ship) error {
	s.Y += i.dist
	return nil
}

func (i north) ExecWaypoint(w *waypoint.Waypoint) error {
	w.Y += i.dist
	return nil
}

type south struct {
	dist int
}

func (i south) ExecShip(s *ship.Ship) error {
	s.Y -= i.dist
	return nil
}

func (i south) ExecWaypoint(w *waypoint.Waypoint) error {
	w.Y -= i.dist
	return nil
}

type east struct {
	dist int
}

func (i east) ExecShip(s *ship.Ship) error {
	s.X += i.dist
	return nil
}

func (i east) ExecWaypoint(w *waypoint.Waypoint) error {
	w.X += i.dist
	return nil
}

type west struct {
	dist int
}

func (i west) ExecShip(s *ship.Ship) error {
	s.X -= i.dist
	return nil
}

func (i west) ExecWaypoint(w *waypoint.Waypoint) error {
	w.X -= i.dist
	return nil
}

type left struct {
	degrees int
}

func (i left) ExecShip(s *ship.Ship) error {
	s.Rotate(-i.degrees)
	return nil
}

func (i left) ExecWaypoint(w *waypoint.Waypoint) error {
	rotations := i.degrees / 90
	for j := 0; j < rotations; j++ {
		wX := w.X
		wY := w.Y
		w.X = -wY
		w.Y = wX
	}

	return nil
}

type right struct {
	degrees int
}

func (i right) ExecShip(s *ship.Ship) error {
	s.Rotate(i.degrees)
	return nil
}

func (i right) ExecWaypoint(w *waypoint.Waypoint) error {
	rotations := i.degrees / 90
	for j := 0; j < rotations; j++ {
		wX := w.X
		wY := w.Y
		w.X = wY
		w.Y = -wX
	}

	return nil
}

type forward struct {
	dist int
}

func (i forward) ExecShip(s *ship.Ship) error {
	if s.Rotation == 0 {
		north{i.dist}.ExecShip(s)
	} else if s.Rotation == 90 {
		east{i.dist}.ExecShip(s)
	} else if s.Rotation == 180 {
		south{i.dist}.ExecShip(s)
	} else if s.Rotation == 270 {
		west{i.dist}.ExecShip(s)
	} else {
		return fmt.Errorf("Invalid ship rotation when moving forward: %d", s.Rotation)
	}

	return nil
}

func (i forward) ExecWaypoint(w *waypoint.Waypoint) error {
	xDelta := w.X * i.dist
	yDelta := w.Y * i.dist

	w.Ship.X += xDelta
	w.Ship.Y += yDelta

	return nil
}
