package instructions

import (
	"fmt"
	"strconv"
)

func Parse(s string) (Instruction, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("cannot parse instruction (input too short): %v", s)
	}

	t := s[0]
	val, err := strconv.Atoi(s[1:])

	if err != nil {
		return nil, fmt.Errorf("cannot parse instruction: %w", err)
	}

	switch t {
	case 'N':
		return north{val}, nil
	case 'S':
		return south{val}, nil
	case 'E':
		return east{val}, nil
	case 'W':
		return west{val}, nil
	case 'L':
		return left{val}, nil
	case 'R':
		return right{val}, nil
	case 'F':
		return forward{val}, nil
	default:
		return nil, fmt.Errorf("cannot parse instruction: unknown instruction type %#v", t)
	}
}
