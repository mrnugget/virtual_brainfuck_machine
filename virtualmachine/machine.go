package main

import "io"

type Machine struct {
	code []*Instruction
	ip   int

	memory [30000]int
	dp     int

	input  io.Reader
	output io.Writer

	readBuf []byte
}

func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins.Type {
		case Plus:
			m.memory[m.dp] += ins.Argument
		case Minus:
			m.memory[m.dp] -= ins.Argument
		case Right:
			m.dp += ins.Argument
		case Left:
			m.dp -= ins.Argument
		case PutChar:
			for i := 0; i < ins.Argument; i++ {
				m.putChar()
			}
		case ReadChar:
			for i := 0; i < ins.Argument; i++ {
				m.readChar()
			}
		case JumpIfZero:
			if m.memory[m.dp] == 0 {
				m.ip = ins.Argument
				continue
			}
		case JumpIfNotZero:
			if m.memory[m.dp] != 0 {
				m.ip = ins.Argument
				continue
			}
		}

		m.ip++
	}
}

func (m *Machine) readChar() {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.readBuf[0])
}

func (m *Machine) putChar() {
	m.readBuf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}
