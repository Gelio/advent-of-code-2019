package statemachine

import "aoc-2020/cmd/19/rule"

type StateMachine struct {
	initialState, finalState int
}

func (m *StateMachine) Matches(line string) bool {
	return false
}

func New(rules []rule.Rule) (StateMachine, error) {
	m := StateMachine{
		finalState: 1,
	}

	// for _, rule := range rules {

	// }

	return m, nil
}
