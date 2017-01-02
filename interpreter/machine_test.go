package main

import (
	"bytes"
	"testing"
)

func TestIncrement(t *testing.T) {
	m := NewMachine("+++++", new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("cell not correctly incremented. got=%d", m.memory[0])
	}
}

func TestDecrement(t *testing.T) {
	input := "++++++++++-----"
	m := NewMachine(input, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("cell not correctly decremented. got=%d", m.memory[0])
	}
}

func TestIncrementingDataPointer(t *testing.T) {
	m := NewMachine("+>++>+++", new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []int{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestDecrementDataPointer(t *testing.T) {
	m := NewMachine(">>+++<++<+", new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []int{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestReadChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)

	m := NewMachine(",>,>,>,>,>,>", in, out)

	m.Execute()

	expectedMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, expected := range expectedMemory {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestPutChar(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	m := NewMachine(".>.>.>.>.>.>", in, out)

	setupMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, value := range setupMemory {
		m.memory[i] = value
	}

	m.Execute()

	output := out.String()
	if output != "ABCDEF" {
		t.Errorf("output wrong. got=%q", output)
	}

}

const HelloWorld = `++++++++[>++++[>++>+++>+++>+<<<<-]>+> +>->>+[<]<-]>>.>---.+++++++ ..+ ++.>>.<-.<.+++.------.--------.>>+.>++.`

func TestHelloWorld(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	m := NewMachine(HelloWorld, in, out)

	m.Execute()

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("output wrong. got=%q", output)
	}
}
