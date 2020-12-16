package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type fieldRange struct {
	from, to int
}

func (r fieldRange) Has(x int) bool {
	return x >= r.from && x <= r.to
}

type ticketFieldSpec struct {
	name   string
	ranges []fieldRange
}

func NewTicketFieldSpec(line string) (ticketFieldSpec, error) {
	spec := ticketFieldSpec{}
	r, err := regexp.Compile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)
	if err != nil {
		return spec, err
	}

	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return spec, fmt.Errorf("cannot match line %s", line)
	}

	spec.name = matches[1]
	firstRange, err := newFieldRange(matches[2], matches[3])
	if err != nil {
		return spec, fmt.Errorf("cannot parse first range: %w", err)
	}

	secondRange, err := newFieldRange(matches[4], matches[5])
	if err != nil {
		return spec, fmt.Errorf("cannot parse second range: %w", err)
	}

	spec.ranges = []fieldRange{firstRange, secondRange}
	return spec, nil
}

func newFieldRange(from, to string) (fieldRange, error) {
	r := fieldRange{}

	fromParsed, err := strconv.Atoi(from)
	if err != nil {
		return r, fmt.Errorf("cannot match \"from\" in range %s: %w", from, err)
	}

	r.from = fromParsed

	toParsed, err := strconv.Atoi(to)
	if err != nil {
		return r, fmt.Errorf("cannot match \"to\" in range %s: %w", to, err)
	}

	r.to = toParsed

	return r, nil
}
