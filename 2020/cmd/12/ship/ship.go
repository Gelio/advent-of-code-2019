package ship

type Ship struct {
	X, Y int
	// 0-359, clockwise, 0 is north
	Rotation int
}

func New() Ship {
	return Ship{Rotation: 90}
}

func (s *Ship) Rotate(degrees int) {
	s.Rotation += (degrees % 360)

	s.normalizeRotation()
}

func (s *Ship) normalizeRotation() {
	if s.Rotation < 0 {
		s.Rotation += 360
	}
	if s.Rotation >= 360 {
		s.Rotation %= 360
	}
}
