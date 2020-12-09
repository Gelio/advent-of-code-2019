package main

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

	visitedInstructions := make(map[int]bool)

	for !visitedInstructions[p.ctx.ip] {
		visitedInstructions[p.ctx.ip] = true
		err = p.Step()
		if err != nil {
			return
		}
	}

	res = p.ctx.acc
	return
}
