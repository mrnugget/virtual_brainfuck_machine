# Brainfuck Interpreter & Virtual Machine

## Interpreter

The `./interpreter` folder contains a Brainfuck interpreter, that reads in a
Brainfuck file and interprets it on the fly.

## Virtual Machine

The `./virtualmachine` folder contains a Brainfuck compiler, that compiles a
Brainfuck file into a custom instruction set. This instruction set is then
executed by a virtual Brainfuck machine.

## Running the code

In each folder you can do the following:

```
go build -o machine && ./machine my_brainfuck_file.b
```

## Tests

Both implementations come with tests. Running all tests just requires this
command:

```
go test ./...
```

## License

MIT. See the LICENSE file.
