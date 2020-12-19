package rule

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	ID      int
	Content interface{}
}

type Alternative struct {
	Options []interface{}
}

type Sequence struct {
	RuleIDs []int
}

type Literal struct {
	Letter rune
}

func ParseLines(lines []string) ([]Rule, error) {
	rules := make(map[int]Rule)
	var ruleIDs []int

	for i, l := range lines {
		rule, err := Parse(l)
		if err != nil {
			return nil, fmt.Errorf("cannot parse rule %d: %w", i, err)
		}

		rules[rule.ID] = rule
		ruleIDs = append(ruleIDs, rule.ID)
	}

	sort.Ints(ruleIDs)

	rulesArr := make([]Rule, ruleIDs[len(ruleIDs)-1]+1)
	for _, ruleID := range ruleIDs {
		rulesArr[ruleID] = rules[ruleID]
	}

	return rulesArr, nil
}

func Parse(line string) (Rule, error) {
	r, err := regexp.Compile(`^(\d+): (.*)$`)
	if err != nil {
		return Rule{}, fmt.Errorf("cannot compile regexp: %w", err)
	}

	matches := r.FindStringSubmatch(line)
	if len(matches) != 3 {
		return Rule{}, fmt.Errorf("cannot match rule %s", line)
	}

	var rule Rule
	rule.ID, err = strconv.Atoi(matches[1])
	if err != nil {
		return Rule{}, fmt.Errorf("cannot parse rule ID %s: %w", matches[1], err)
	}

	parts := strings.Split(matches[2], " | ")
	if len(parts) == 1 {
		rule.Content, err = parsePart(parts[0])
		if err != nil {
			return rule, fmt.Errorf("cannot parse part %s: %w", parts[0], err)
		}
	} else {
		var options []interface{}
		for _, part := range parts {
			option, err := parsePart(part)
			if err != nil {
				return rule, fmt.Errorf("cannot parse part %s: %w", part, err)
			}
			options = append(options, option)
		}

		rule.Content = Alternative{options}
	}

	return rule, nil
}

var litRegexp = regexp.MustCompile(`"(\w)"`)

func parsePart(part string) (interface{}, error) {
	matches := litRegexp.FindStringSubmatch(part)
	if len(matches) == 2 {
		return Literal{getFirstRune(matches[1])}, nil
	}

	var seq Sequence

	rawRuleIDs := strings.Split(part, " ")
	for _, rawID := range rawRuleIDs {
		ruleID, err := strconv.Atoi(rawID)
		if err != nil {
			return seq, fmt.Errorf("cannot parse number %s: %w", rawID, err)
		}
		seq.RuleIDs = append(seq.RuleIDs, ruleID)
	}

	return seq, nil
}

func getFirstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}
