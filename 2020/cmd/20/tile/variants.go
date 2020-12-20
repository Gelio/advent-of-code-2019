package tile

type Rotation int

const (
	NoRotation Rotation = iota
	Rotate90
	Rotate180
	Rotate270
)

type Flip int

const (
	NoFlip Flip = iota
	Vertical
	Horizontal
)

type Borders struct {
	Bottom string
	Top    string
	Right  string
	Left   string
}

func (base Tile) GetAllVariants() []Tile {
	variants := make([]Tile, 0, 4*2+1)
	variants = append(variants, base)

	encounteredBorders := map[Borders]bool{base.Borders: true}

	t := base
	for r := NoRotation; r <= Rotate270; r++ {
		t.Borders.rotate90()
		t.Rotation = r

		// No flip
		if !encounteredBorders[t.Borders] {
			encounteredBorders[t.Borders] = true
			variants = append(variants, t)
		}

		originalBorders := t.Borders

		// Vertical flip
		t.Borders.flipVertical()
		t.Flip = Vertical

		if !encounteredBorders[t.Borders] {
			encounteredBorders[t.Borders] = true
			variants = append(variants, t)
		}

		// Horizontal flip
		t.Borders = originalBorders
		t.Borders.flipHorizontal()
		t.Flip = Horizontal
		if !encounteredBorders[t.Borders] {
			encounteredBorders[t.Borders] = true
			variants = append(variants, t)
		}

		// Restore flip for next iterations
		t.Borders = originalBorders
		t.Flip = NoFlip
	}

	return variants
}

func (b *Borders) rotate90() {
	top := b.Top
	b.Top = b.Left
	b.Left = b.Bottom
	b.Bottom = b.Right
	b.Right = top
}

func (b *Borders) flipVertical() {
	b.Top, b.Bottom = b.Bottom, b.Top
	b.Left = swapFirstWithLast(b.Left)
	b.Right = swapFirstWithLast(b.Right)
}

func (b *Borders) flipHorizontal() {
	b.Left, b.Right = b.Right, b.Left
	b.Top = swapFirstWithLast(b.Top)
	b.Bottom = swapFirstWithLast(b.Bottom)
}

func swapFirstWithLast(s string) string {
	return s[len(s)-1:] + s[1:len(s)-1] + s[0:1]
}
