package interpreter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/kriticalflare/go-brainfuck-yourself/token"
)

type Interpreter struct {
	Memory  []uint8
	ptr     int
	program []token.Token
	reader  *bufio.Reader
}

func New(program []token.Token, memorySize uint) *Interpreter {
	reader := bufio.NewReader(os.Stdin)
	return &Interpreter{Memory: make([]uint8, memorySize), ptr: 0, program: program, reader: reader}
}

func (i *Interpreter) Run() {
	// fmt.Printf("memory: %v\n", i.Memory)
	commandIdx := 0
	for {
		if commandIdx >= len(i.program) {
			break
		}
		tok := i.program[commandIdx]

		switch t := tok.(type) {
		case *token.Add:
			{
				i.Memory[i.ptr]++
			}
		case *token.Subtract:
			{
				i.Memory[i.ptr]--
			}
		case *token.Next:
			{
				i.ptr++
			}
		case *token.Prev:
			{
				i.ptr--
			}
		case *token.JumpZero:
			{
				if i.Memory[i.ptr] == 0 {
					commandIdx = int(t.Target)
				}
			}
		case *token.JumpNonZero:
			{
				if i.Memory[i.ptr] != 0 {
					commandIdx = int(t.Target)
				}
			}
		case *token.Output:
			{
				fmt.Printf("%c", i.Memory[i.ptr])
			}
		case *token.Input:
			{
				// fmt.Printf("waiting for input %v\n", commandIdx)
				input, err := i.reader.ReadByte()
				if errors.Is(err, io.EOF) {
					fmt.Printf("EOF\n")
					i.Memory[i.ptr] = 0
					continue
				}
				if err != nil {
					fmt.Printf("failed to read input:%d. Error: %v\n", t.SourcePosition(), err)
				}
				if input == 13 {
					// skip carriage return in windows
					input, err = i.reader.ReadByte()
				}
				if err != nil {
					fmt.Printf("failed to read input:%d. Error: %v\n", t.SourcePosition(), err)
				}
				// fmt.Printf("read '%d'\n", input)
				i.Memory[i.ptr] = input
			}
		}
		commandIdx++
	}
}
