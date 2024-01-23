package encapsulator

import "github.com/Tch1b0/brainfuck-interpreter/internal/token"

const (
	ACT_INSTR = iota
	ACT_BLOCK
)

type Block struct {
	Actions []Action
}

type Action struct {
	Type           int
	InstructionTok int
	Block          *Block
}

func BuildBlock(tokens []int) (Block, int) {
	block := Block{
		Actions: []Action{},
	}

	idx := 0

	for idx < len(tokens) {
		tok := tokens[idx]
		act := Action{}

		if tok == token.BEGIN_LOOP {
			b, nIdx := BuildBlock(tokens[idx+1:])
			act.Type = ACT_BLOCK
			act.Block = &b
			idx += nIdx + 1
		} else if tok == token.END_LOOP {
			return block, idx
		} else {
			act.Type = ACT_INSTR
			act.InstructionTok = tok
		}

		block.Actions = append(block.Actions, act)
		idx++
	}

	return block, idx
}
