package main

import (
	"flag"
	"fmt"
	"lem-in/solver"
	"log"
	"os"
)

const setBold = "\033[1m"
const reset = "\033[0m"

const HELP = `Usage: ` + setBold + `lem-in [file.txt]` + reset + ` to read farm configuration from file and solve it
Example: ` + setBold + `lem-in ./examples/example00.txt` + reset

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println(HELP)
	}
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Invalid arguments")
		flag.Usage()
		return
	}

	cfgBytes, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	graph, err := solver.ReadGraph(string(cfgBytes))
	if err != nil {
		log.Fatalf("ERROR: invalid data format, %v", err)
	}

	err = graph.FindPaths()
	if err != nil {
		log.Fatalf("ERROR: invalid data format, %v", err)
	}

	fmt.Printf("%s\n\n", cfgBytes)

	solution := graph.Solve()

	//for i, path := range solution {
	//	for _, point := range *path {
	//		fmt.Print(point.Name, " ")
	//	}
	//	fmt.Printf(" (%v)\n", &solution[i])
	//}

	fmt.Println(graph.FormatSolution(solution))
}
