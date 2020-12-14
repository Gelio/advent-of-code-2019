package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type setMemory struct {
	memIndex int
	value    int
}

type setMask struct {
	mask string
}

func parseOpLine(line string) (interface{}, error) {
	setMaskRegexp, err := regexp.Compile("^mask = ([01X]+)$")

	if err != nil {
		return nil, fmt.Errorf("cannot compile setMaskRegexp: %w", err)
	}

	matches := setMaskRegexp.FindStringSubmatch(line)

	if len(matches) > 0 {
		rawMask := matches[1]
		return setMask{rawMask}, nil
	}

	setMemoryRegexp, err := regexp.Compile(`^mem\[(\d+)\] = (\d+)$`)
	if err != nil {
		return nil, fmt.Errorf("cannot compile setMemoryRegexp: %w", err)
	}

	matches = setMemoryRegexp.FindStringSubmatch(line)

	if len(matches) > 0 {
		memIndex, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("cannot parse memory index %s", matches[1])
		}

		value, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("cannot parse memory value %s", matches[2])
		}

		return setMemory{memIndex, value}, nil
	}

	return nil, fmt.Errorf("cannot parse line %s", line)
}
