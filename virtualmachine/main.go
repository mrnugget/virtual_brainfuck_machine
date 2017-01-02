package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// var cpuprofile string
// var memprofile string
//
// func init() {
// 	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
// 	flag.StringVar(&memprofile, "memprofile", "", "write mem profile to file")
// }

func main() {
	// flag.Parse()
	// args := flag.Args()
	//
	// if len(args) != 1 {
	// 	fmt.Fprintf(os.Stderr, "no filename given\n")
	// 	os.Exit(-1)
	// }
	//
	// if cpuprofile != "" {
	// 	f, err := os.Create(cpuprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }

	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	m := NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
	//
	// if memprofile != "" {
	// 	f, err := os.Create(memprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.WriteHeapProfile(f)
	// 	f.Close()
	// 	return
	// }
}
