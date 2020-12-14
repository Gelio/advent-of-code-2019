package main

// maskA applies the bitmasking rules from part A
type maskA struct {
	orMask, andMask uint
}

func newMaskA(line string) maskA {
	m := maskA{}

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

func (m maskA) Apply(value int) int {
	return int(uint(value)&m.andMask | m.orMask)
}
