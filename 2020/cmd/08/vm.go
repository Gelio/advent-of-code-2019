package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Context struct {
	ip, acc int
}

func (c *Context) nextInstr() {
	c.ip++
}

type Instruction interface {
	Exec(c *Context)
}

type accInstr struct {
	v int
}

// QUESTION: is `*accInstr` ok, or I would be better of with `accInstr`? What's the difference?
func (i *accInstr) Exec(c *Context) {
	c.acc += i.v
	c.nextInstr()
}

type nopInstr struct {
	v int
}

func (i *nopInstr) Exec(c *Context) {
	c.nextInstr()
}

type jmpInstr struct {
	v int
}

func (i *jmpInstr) Exec(c *Context) {
	c.ip += i.v
}

func parseInstruction(s string) (instr Instruction, err error) {
	r := regexp.MustCompile(`^(acc|nop|jmp) ((\+|-)\d+)$`)

	matches := r.FindStringSubmatch(s)

	if matches == nil {
		err = fmt.Errorf("Invalid instruction: %v", s)
		return
	}

	var num int
	num, err = strconv.Atoi(matches[2])

	if err != nil {
		err = fmt.Errorf("Cannot parse number in instruction %#v: %w", s, err)
		return
	}

	switch matches[1] {
	case "nop":
		// QUESTION: this one is weird to me. It does not compile if I expect `*Instruction` in params.
		// Why?
		// NOTE: cannot use &(nopInstr literal) (value of type *nopInstr) as *Instruction value in assignment
		instr = &nopInstr{v: num}
		return

	case "acc":
		instr = &accInstr{v: num}
		return

	case "jmp":
		instr = &jmpInstr{v: num}
		return

	default:
		err = fmt.Errorf("Invalid instruction keyword %#v in instruction %#v", matches[1], s)
		return
	}
}

type Program struct {
	// QUESTION: is `[]Instruction` ok, or should I do `[]*Instruction`?
	instructions []Instruction
	ctx          *Context
}

// QUESTIONS: are such factory methods common?
func NewProgram(instrs []Instruction) *Program {
	return &Program{
		instructions: instrs,
		ctx: &Context{
			ip:  0,
			acc: 0,
		},
	}
}

func (p *Program) Step() error {
	if p.ctx.ip < 0 {
		return fmt.Errorf("Negative instruction pointer (%d)", p.ctx.ip)
	}

	if p.ctx.ip >= len(p.instructions) {
		return fmt.Errorf("IP is past the list of instructions (%d, list of instructions has length %d)", p.ctx.ip, len(p.instructions))
	}

	instr := p.instructions[p.ctx.ip]

	if instr == nil {
		return fmt.Errorf("Invalid instruction pointer %d", p.ctx.ip)
	}

	instr.Exec(p.ctx)

	return nil
}

func (p *Program) hasFinished() bool {
	return p.ctx.ip == len(p.instructions)
}

func (p *Program) Run() error {
	visitedInstructions := make(map[int]bool)

	for !visitedInstructions[p.ctx.ip] && !p.hasFinished() {
		visitedInstructions[p.ctx.ip] = true

		if err := p.Step(); err != nil {
			return err
		}
	}

	return nil
}
