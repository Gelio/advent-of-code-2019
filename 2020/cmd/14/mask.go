package main

// maskA applies the bitmasking rules from part A
type maskA struct {
	orMask, andMask uint
}

func newNoopMaskA(line string) maskA {
	m := maskA{}

	var oneBit uint = 1

	for range line {
		m.andMask |= oneBit
		oneBit <<= 1
	}

	return m
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

// masksB is used to generate all masks in part B
type masksB []maskA

func newMaskB(line string) masksB {
	var m masksB
	var oneBit uint = 1 << (len(line) - 1)
	currentMask := newNoopMaskA(line)

	m.generateBitmasksRecursive(line, oneBit, currentMask)

	return m
}

func (m *masksB) generateBitmasksRecursive(line string, oneBit uint, currentMask maskA) {
	for i, c := range line {
		nextOneBit := oneBit >> 1

		switch c {
		case '1':
			currentMask.orMask |= oneBit

		case '0':

		case 'X':
			// When X becomes 0
			m.generateBitmasksRecursive(line[i+1:], nextOneBit, maskA{
				orMask:  currentMask.orMask,
				andMask: currentMask.andMask - oneBit,
			})

			// When X becomes 1
			m.generateBitmasksRecursive(line[i+1:], nextOneBit, maskA{
				orMask:  currentMask.orMask + oneBit,
				andMask: currentMask.andMask,
			})
			return
		}

		oneBit = nextOneBit
	}

	*m = append(*m, currentMask)
}

func (m masksB) Apply(value int) []int {
	var res []int

	for _, mask := range m {
		res = append(res, mask.Apply(value))
	}

	return res
}
