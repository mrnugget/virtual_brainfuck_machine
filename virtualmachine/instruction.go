package main

type InsType byte

const (
	Plus          InsType = '+'
	Minus         InsType = '-'
	Right         InsType = '>'
	Left          InsType = '<'
	PutChar       InsType = '.'
	ReadChar      InsType = ','
	JumpIfZero    InsType = '['
	JumpIfNotZero InsType = ']'
)

type Instruction struct {
	Type     InsType
	Argument int
}
