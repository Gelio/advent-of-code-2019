package statemachine

import (
	"aoc-2020/cmd/19/rule"
	"errors"
	"fmt"
)

type state int
type move struct {
	state
	letter rune
}

type StateMachine struct {
	initialState, finalState state
	nextEmptyState           state

	graph map[move]state
	rules []rule.Rule
}

func New(rules []rule.Rule) (StateMachine, error) {
	m := StateMachine{
		finalState:     1,
		graph:          make(map[move]state),
		nextEmptyState: 2,
		rules:          rules,
	}

	if err := m.processContent(rules[0].Content, m.initialState, &m.finalState); err != nil {
		return m, err
	}

	return m, nil
}

func (m *StateMachine) processContent(content interface{}, fromState state, toState *state) error {
	switch c := content.(type) {
	case rule.Literal:
		m.processLiteral(c, fromState, toState)
		return nil

	case rule.Alternative:
		return m.processAlternative(c, fromState, toState)

	case rule.Sequence:
		return m.processSequence(c, fromState, toState)

	default:
		return fmt.Errorf("invalid content %v", content)
	}
}

func (m *StateMachine) processLiteral(r rule.Literal, fromState state, toState *state) {
	if s, ok := m.graph[move{fromState, r.Letter}]; ok {
		*toState = s
	} else {
		m.graph[move{fromState, r.Letter}] = *toState
		fmt.Println(move{fromState, r.Letter}, *toState)
	}
}

func (m *StateMachine) processSequence(r rule.Sequence, fromState state, toState *state) error {
	if len(r.RuleIDs) == 0 {
		return errors.New("invalid sequence of 0 elements")
	}

	lastIDIndex := len(r.RuleIDs) - 1
	prevState := fromState
	for i, id := range r.RuleIDs[:lastIDIndex] {
		nextState := m.nextEmptyState
		m.nextEmptyState++
		nextEmptyState := m.nextEmptyState

		if err := m.processContent(m.rules[id].Content, prevState, &nextState); err != nil {
			return fmt.Errorf("cannot process item %d (rule %d) in sequence: %w", i+1, id, err)
		}

		if unusedNewState := nextState == nextEmptyState; unusedNewState {
			m.nextEmptyState--
		}

		prevState = nextState
	}

	if err := m.processContent(m.rules[r.RuleIDs[lastIDIndex]].Content, prevState, toState); err != nil {
		return fmt.Errorf("cannot process last item (rule %d) in sequence: %w", r.RuleIDs[lastIDIndex], err)
	}

	return nil
}

func (m *StateMachine) processAlternative(r rule.Alternative, fromState state, toState *state) error {
	for i, option := range r.Options {
		if err := m.processContent(option, fromState, toState); err != nil {
			return fmt.Errorf("cannot process option %d %v: %w", i, option, err)
		}
	}

	return nil
}

func (m *StateMachine) Matches(line string) bool {
	state := m.initialState

	for _, r := range line {
		var ok bool
		if state, ok = m.graph[move{state, r}]; !ok {
			return false
		}
	}

	return state == m.finalState
}
