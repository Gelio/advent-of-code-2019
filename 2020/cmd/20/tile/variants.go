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

type variantGenerator struct {
	variants           []Tile
	encounteredBorders map[Borders]bool
}

func (vg *variantGenerator) addTile(t Tile) {
	if !vg.encounteredBorders[t.Borders] {
		vg.encounteredBorders[t.Borders] = true
		vg.variants = append(vg.variants, t)
	}
}

func (t Tile) GetAllVariants() []Tile {
	var vg variantGenerator
	vg.variants = make([]Tile, 0, 8)
	vg.encounteredBorders = make(map[Borders]bool)

	t1 := t
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)

	t1 = t
	t1.Borders.flipVertical()
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)
	t1.Borders.rotate90()
	vg.addTile(t1)

	return vg.variants
}

func (b *Borders) rotate90() {
	top := b.Top
	b.Top = reverseString(b.Left)
	b.Left = reverseString(b.Bottom)
	b.Bottom = reverseString(b.Right)
	b.Right = top
}

func (b *Borders) flipVertical() {
	b.Top, b.Bottom = b.Bottom, b.Top
	b.Left = swapFirstWithLast(b.Left)
	b.Right = swapFirstWithLast(b.Right)
}

func swapFirstWithLast(s string) string {
	return s[len(s)-1:] + s[1:len(s)-1] + s[0:1]
}

func reverseString(s string) string {
	res := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	return string(res)
}
