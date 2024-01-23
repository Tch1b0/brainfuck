package interpreter

import (
	"fmt"
	"os"

	"github.com/Tch1b0/brainfuck-interpreter/internal/encapsulator"
	"github.com/Tch1b0/brainfuck-interpreter/internal/token"
)

var (
	memory = [30_000]byte{}
	DP     = 0
)

func Execute(block encapsulator.Block, topLevel bool) {
	for {
		for _, act := range block.Actions {
			if act.Type == encapsulator.ACT_BLOCK {
				Execute(*act.Block, false)
				continue
			}

			switch act.InstructionTok {
			case token.INC_DP:
				DP++
			case token.DEC_DP:
				DP--
			case token.INC_VAL:
				memory[DP]++
			case token.DEC_VAL:
				memory[DP]--
			case token.OUT:
				fmt.Print(string(rune(memory[DP])))
			case token.IN:
				memory[DP] = readInput()
			}

		}

		if topLevel || memory[DP] == 0 {
			break
		}
	}
}

func readInput() byte {
	b := make([]byte, 1)
	os.Stdin.Read(b)
	return b[0]
}
