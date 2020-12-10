package main

import "fmt"

func SolveA(lines []string) (res int, err error) {
	p := NewProgram([]Instruction{})

	for _, line := range lines {
		var instr Instruction
		instr, err = parseInstruction(line)
		if err != nil {
			return
		}

		p.instructions = append(p.instructions, instr)
	}

	err = p.Run()
	if err != nil {
		return
	}

	res = p.ctx.acc
	return
}

func SolveB(lines []string) (res int, err error) {
	initialInstrs := make([]Instruction, 0, len(lines))

	for _, line := range lines {
		var instr Instruction
		instr, err = parseInstruction(line)
		if err != nil {
			return
		}

		initialInstrs = append(initialInstrs, instr)
	}

	resChan := make(chan int)

	for idx, instr := range initialInstrs {
		switch i := instr.(type) {
		case *accInstr:
			continue

		case *jmpInstr:
			// NOTE: I want each goroutine to work on a `initialInstrs` array, but with 1 element changed.
			// QUESTION: how should I do it?
			// NOTE: see `runWithChangedInstructions` for my idea on how I have done it
			go runWithChangedInstructions(initialInstrs, idx, &nopInstr{v: i.v}, resChan)

		case *nopInstr:
			go runWithChangedInstructions(initialInstrs, idx, &jmpInstr{v: i.v}, resChan)

		default:
			err = fmt.Errorf("Unknown instruction type %T", instr)
			return
		}
	}

	res = <-resChan
	return
}

func runWithChangedInstructions(instrs []Instruction, idx int, newInstr Instruction, resChan chan int) {
	newInstrs := make([]Instruction, len(instrs), len(instrs))
	copy(newInstrs, instrs)
	newInstrs[idx] = newInstr
	p := NewProgram(newInstrs)

	err := p.Run()
	if err == nil && p.hasFinished() {
		resChan <- p.ctx.acc
	}
}
