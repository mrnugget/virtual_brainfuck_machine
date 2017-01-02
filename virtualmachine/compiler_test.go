package main

import "testing"

func TestCompile(t *testing.T) {
	input := `
	+++++
	-----
	+++++
	>>>>>
	<<<<<
	`
	expected := []*Instruction{
		&Instruction{Plus, 5},
		&Instruction{Minus, 5},
		&Instruction{Plus, 5},
		&Instruction{Right, 5},
		&Instruction{Left, 5},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []*Instruction{
		&Instruction{Plus, 1},
		&Instruction{JumpIfZero, 7},
		&Instruction{Plus, 1},
		&Instruction{JumpIfZero, 5},
		&Instruction{Plus, 1},
		&Instruction{JumpIfNotZero, 3},
		&Instruction{Plus, 1},
		&Instruction{JumpIfNotZero, 1},
		&Instruction{Plus, 1},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileEverything(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []*Instruction{
		&Instruction{Plus, 3},
		&Instruction{JumpIfZero, 7},
		&Instruction{Minus, 3},
		&Instruction{JumpIfZero, 5},
		&Instruction{Plus, 1},
		&Instruction{JumpIfNotZero, 3},
		&Instruction{Right, 3},
		&Instruction{JumpIfNotZero, 1},
		&Instruction{Left, 3},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}
