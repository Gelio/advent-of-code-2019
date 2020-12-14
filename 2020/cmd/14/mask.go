package main

type mask struct {
	orMask, andMask uint
}

func newMask(line string) mask {
	m := mask{}

	maskLen := len(line)

	for i, c := range line {
		bitIndex := maskLen - 1 - i
		switch c {
		case 'X':
			// noop:
			// m.orMask |= (0 << bitIndex)
			m.andMask |= 1 << bitIndex

		case '1':
			m.orMask |= 1 << bitIndex

		case '0':
			// noop
			// m.andMask |= 0 << bitIndex
		}
	}

	return m
}

func (m mask) Apply(value int) int {
	return int(uint(value)&m.andMask | m.orMask)
}
