package solver

import (
	. "lem-in/solver"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func twoEdgeSequencesTheSame(eOne, eTwo []*Edge) bool {

	// edges names sequence generator in form of map keys with value 1
	var eName = func(edges []*Edge) (res map[string]int) {
		res = make(map[string]int, len(edges))
		for _, e := range edges {
			res[e.From.Name+"-"+e.To.Name] = 1
		}
		return
	}
	// convert to map keys sequence for comparing
	one, two := eName(eOne), eName(eTwo)
	for key := range one {
		if _, ok := two[key]; !ok {
			return false
		}
	}
	return true && len(one) == len(two)
}

func TestSolver(t *testing.T) {
	goodTests := []struct {
		fileName       string
		reference      *Graph
		solutionLength int
		edgeLength     int
	}{
		{fileName: "examples/good/example00.txt",
			reference: &Graph{
				AntNum: 4,
				Start:  &Point{Name: "0", X: 0, Y: 3},
				End:    &Point{Name: "1", X: 8, Y: 3},
				Points: []*Point{
					{Name: "2", X: 2, Y: 5},
					{Name: "3", X: 4, Y: 0},
				},
				Edges: []*Edge{
					{From: &Point{Name: "0"}, To: &Point{Name: "2"}},
					{From: &Point{Name: "2"}, To: &Point{Name: "3"}},
					{From: &Point{Name: "3"}, To: &Point{Name: "1"}},
				},
			},
			solutionLength: 6,
			edgeLength:     3,
		},
		// TODO: add Points and Edges to tests
		{fileName: "examples/good/example01.txt",
			reference: &Graph{AntNum: 10,
				Start: &Point{Name: "start", X: 1, Y: 6},
				End:   &Point{Name: "end", X: 11, Y: 6},
				Points: []*Point{
					{Name: "0", X: 4, Y: 8},
					{Name: "o", X: 6, Y: 8},
					{Name: "n", X: 6, Y: 6},
					{Name: "e", X: 8, Y: 4},
					{Name: "t", X: 1, Y: 9},
					{Name: "E", X: 5, Y: 9},
					{Name: "a", X: 8, Y: 9},
					{Name: "m", X: 8, Y: 6},
					{Name: "h", X: 4, Y: 6},
					{Name: "A", X: 5, Y: 2},
					{Name: "c", X: 8, Y: 1},
					{Name: "k", X: 11, Y: 2},
				},
				Edges: []*Edge{
					{From: &Point{Name: "start"}, To: &Point{Name: "t"}},
					{From: &Point{Name: "n"}, To: &Point{Name: "e"}},
					{From: &Point{Name: "a"}, To: &Point{Name: "m"}},
					{From: &Point{Name: "A"}, To: &Point{Name: "c"}},
					{From: &Point{Name: "0"}, To: &Point{Name: "o"}},
					{From: &Point{Name: "E"}, To: &Point{Name: "a"}},
					{From: &Point{Name: "k"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "h"}},
					{From: &Point{Name: "o"}, To: &Point{Name: "n"}},
					{From: &Point{Name: "m"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "t"}, To: &Point{Name: "E"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "0"}},
					{From: &Point{Name: "h"}, To: &Point{Name: "A"}},
					{From: &Point{Name: "e"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "c"}, To: &Point{Name: "k"}},
					{From: &Point{Name: "n"}, To: &Point{Name: "m"}},
					{From: &Point{Name: "h"}, To: &Point{Name: "n"}},
				},
			},
			solutionLength: 8,
			edgeLength:     17,
		},
		{fileName: "examples/good/example02.txt",
			reference: &Graph{AntNum: 20,
				Start: &Point{Name: "0", X: 2, Y: 0},
				End:   &Point{Name: "3", X: 5, Y: 3},
				Points: []*Point{
					{Name: "1", X: 4, Y: 1},
					{Name: "2", X: 6, Y: 0},
				}, Edges: []*Edge{
					{From: &Point{Name: "0"}, To: &Point{Name: "1"}},
					{From: &Point{Name: "0"}, To: &Point{Name: "3"}},
					{From: &Point{Name: "1"}, To: &Point{Name: "2"}},
					{From: &Point{Name: "3"}, To: &Point{Name: "2"}},
				},
			},
			solutionLength: 11,
			edgeLength:     4,
		},
		{fileName: "examples/good/example03.txt",
			reference: &Graph{AntNum: 4,
				Start: &Point{Name: "0", X: 1, Y: 4},
				End:   &Point{Name: "5", X: 6, Y: 4},
				Points: []*Point{
					{Name: "4", X: 5, Y: 4},
					{Name: "1", X: 3, Y: 6},
					{Name: "2", X: 3, Y: 4},
					{Name: "3", X: 3, Y: 1},
				}, Edges: []*Edge{
					{From: &Point{Name: "0"}, To: &Point{Name: "1"}},
					{From: &Point{Name: "2"}, To: &Point{Name: "4"}},
					{From: &Point{Name: "1"}, To: &Point{Name: "4"}},
					{From: &Point{Name: "0"}, To: &Point{Name: "2"}},
					{From: &Point{Name: "4"}, To: &Point{Name: "5"}},
					{From: &Point{Name: "3"}, To: &Point{Name: "0"}},
					{From: &Point{Name: "4"}, To: &Point{Name: "3"}},
				},
			},
			solutionLength: 6,
			edgeLength:     7,
		},
		{fileName: "examples/good/example04.txt",
			reference: &Graph{AntNum: 9,
				Start: &Point{Name: "richard", X: 0, Y: 6},
				End:   &Point{Name: "peter", X: 14, Y: 6},
				Points: []*Point{
					{Name: "gilfoyle", X: 6, Y: 3},
					{Name: "erlich", X: 9, Y: 6},
					{Name: "dinish", X: 6, Y: 9},
					{Name: "jimYoung", X: 11, Y: 7},
				}, Edges: []*Edge{
					{From: &Point{Name: "richard"}, To: &Point{Name: "dinish"}},
					{From: &Point{Name: "dinish"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "gilfoyle"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "peter"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "erlich"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "jimYoung"}, To: &Point{Name: "peter"}},
				},
			},
			solutionLength: 6,
			edgeLength:     8,
		},
		{fileName: "examples/good/example05.txt",
			reference: &Graph{AntNum: 9,
				Start: &Point{Name: "start", X: 0, Y: 3},
				End:   &Point{Name: "end", X: 10, Y: 1},
				Points: []*Point{
					{Name: "C0", X: 1, Y: 0},
					{Name: "C1", X: 2, Y: 0},
					{Name: "C2", X: 3, Y: 0},
					{Name: "C3", X: 4, Y: 0},
					{Name: "I4", X: 5, Y: 0},
					{Name: "I5", X: 6, Y: 0},
					{Name: "A0", X: 1, Y: 2},
					{Name: "A1", X: 2, Y: 1},
					{Name: "A2", X: 4, Y: 1},
					{Name: "B0", X: 1, Y: 4},
					{Name: "B1", X: 2, Y: 4},
					{Name: "E2", X: 6, Y: 4},
					{Name: "D1", X: 6, Y: 3},
					{Name: "D2", X: 7, Y: 3},
					{Name: "D3", X: 8, Y: 3},
					{Name: "H4", X: 4, Y: 2},
					{Name: "H3", X: 5, Y: 2},
					{Name: "F2", X: 6, Y: 2},
					{Name: "F3", X: 7, Y: 2},
					{Name: "F4", X: 8, Y: 2},
					{Name: "G0", X: 1, Y: 5},
					{Name: "G1", X: 2, Y: 5},
					{Name: "G2", X: 3, Y: 5},
					{Name: "G3", X: 4, Y: 5},
					{Name: "G4", X: 6, Y: 5},
				}, Edges: []*Edge{
					{From: &Point{Name: "H3"}, To: &Point{Name: "F2"}},
					{From: &Point{Name: "H3"}, To: &Point{Name: "H4"}},
					{From: &Point{Name: "H4"}, To: &Point{Name: "A2"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "G0"}},
					{From: &Point{Name: "G0"}, To: &Point{Name: "G1"}},
					{From: &Point{Name: "G1"}, To: &Point{Name: "G2"}},
					{From: &Point{Name: "G2"}, To: &Point{Name: "G3"}},
					{From: &Point{Name: "G3"}, To: &Point{Name: "G4"}},
					{From: &Point{Name: "G4"}, To: &Point{Name: "D3"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "A0"}},
					{From: &Point{Name: "A0"}, To: &Point{Name: "A1"}},
					{From: &Point{Name: "A0"}, To: &Point{Name: "D1"}},
					{From: &Point{Name: "A1"}, To: &Point{Name: "A2"}},
					{From: &Point{Name: "A1"}, To: &Point{Name: "B1"}},
					{From: &Point{Name: "A2"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "A2"}, To: &Point{Name: "C3"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "B0"}},
					{From: &Point{Name: "B0"}, To: &Point{Name: "B1"}},
					{From: &Point{Name: "B1"}, To: &Point{Name: "E2"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "C0"}},
					{From: &Point{Name: "C0"}, To: &Point{Name: "C1"}},
					{From: &Point{Name: "C1"}, To: &Point{Name: "C2"}},
					{From: &Point{Name: "C2"}, To: &Point{Name: "C3"}},
					{From: &Point{Name: "C3"}, To: &Point{Name: "I4"}},
					{From: &Point{Name: "D1"}, To: &Point{Name: "D2"}},
					{From: &Point{Name: "D1"}, To: &Point{Name: "F2"}},
					{From: &Point{Name: "D2"}, To: &Point{Name: "E2"}},
					{From: &Point{Name: "D2"}, To: &Point{Name: "D3"}},
					{From: &Point{Name: "D2"}, To: &Point{Name: "F3"}},
					{From: &Point{Name: "D3"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "F2"}, To: &Point{Name: "F3"}},
					{From: &Point{Name: "F3"}, To: &Point{Name: "F4"}},
					{From: &Point{Name: "F4"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "I4"}, To: &Point{Name: "I5"}},
					{From: &Point{Name: "I5"}, To: &Point{Name: "end"}},
				},
			},
			solutionLength: 8,
			edgeLength:     35,
		},
		{fileName: "examples/good/example06.txt",
			reference: &Graph{AntNum: 100,
				Start: &Point{Name: "richard", X: 0, Y: 6},
				End:   &Point{Name: "peter", X: 14, Y: 6},
				Points: []*Point{
					{Name: "gilfoyle", X: 6, Y: 3},
					{Name: "erlich", X: 9, Y: 6},
					{Name: "dinish", X: 6, Y: 9},
					{Name: "jimYoung", X: 11, Y: 7},
				}, Edges: []*Edge{
					{From: &Point{Name: "richard"}, To: &Point{Name: "dinish"}},
					{From: &Point{Name: "dinish"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "gilfoyle"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "peter"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "erlich"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "jimYoung"}, To: &Point{Name: "peter"}},
				},
			},
			solutionLength: 52,
			edgeLength:     8,
		},
		{fileName: "examples/good/example07.txt",
			reference: &Graph{AntNum: 1000,
				Start: &Point{Name: "richard", X: 0, Y: 6},
				End:   &Point{Name: "peter", X: 14, Y: 6},
				Points: []*Point{
					{Name: "gilfoyle", X: 6, Y: 3},
					{Name: "erlich", X: 9, Y: 6},
					{Name: "dinish", X: 6, Y: 9},
					{Name: "jimYoung", X: 11, Y: 7},
				}, Edges: []*Edge{
					{From: &Point{Name: "richard"}, To: &Point{Name: "dinish"}},
					{From: &Point{Name: "dinish"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "gilfoyle"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "peter"}},
					{From: &Point{Name: "gilfoyle"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "richard"}, To: &Point{Name: "erlich"}},
					{From: &Point{Name: "erlich"}, To: &Point{Name: "jimYoung"}},
					{From: &Point{Name: "jimYoung"}, To: &Point{Name: "peter"}},
				},
			},
			solutionLength: 502,
			edgeLength:     8,
		},
		{fileName: "examples/good/hard.txt",
			reference: &Graph{AntNum: 5,
				Start: &Point{Name: "start", X: 0, Y: 2},
				End:   &Point{Name: "end", X: 5, Y: 2},
				Points: []*Point{
					{Name: "1", X: 0, Y: 0},
					{Name: "2", X: 1, Y: 1},
					{Name: "3", X: 2, Y: 0},
					{Name: "4", X: 2, Y: 2},
					{Name: "5", X: 3, Y: 2},
					{Name: "6", X: 3, Y: 1},
					{Name: "7", X: 4, Y: 0},
					{Name: "8", X: 4, Y: 2},
					{Name: "10", X: 3, Y: 3},
					{Name: "11", X: 4, Y: 4},
					{Name: "12", X: 5, Y: 4},
					{Name: "13", X: 4, Y: 3},
				}, Edges: []*Edge{
					{From: &Point{Name: "start"}, To: &Point{Name: "1"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "2"}},
					{From: &Point{Name: "start"}, To: &Point{Name: "4"}},
					{From: &Point{Name: "1"}, To: &Point{Name: "3"}},
					{From: &Point{Name: "3"}, To: &Point{Name: "7"}},
					{From: &Point{Name: "2"}, To: &Point{Name: "6"}},
					{From: &Point{Name: "7"}, To: &Point{Name: "8"}},
					{From: &Point{Name: "6"}, To: &Point{Name: "5"}},
					{From: &Point{Name: "4"}, To: &Point{Name: "5"}},
					{From: &Point{Name: "5"}, To: &Point{Name: "8"}},
					{From: &Point{Name: "5"}, To: &Point{Name: "13"}},
					{From: &Point{Name: "13"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "4"}, To: &Point{Name: "10"}},
					{From: &Point{Name: "10"}, To: &Point{Name: "11"}},
					{From: &Point{Name: "11"}, To: &Point{Name: "12"}},
					{From: &Point{Name: "12"}, To: &Point{Name: "end"}},
					{From: &Point{Name: "8"}, To: &Point{Name: "end"}},
				},
			},
			solutionLength: 6,
			edgeLength:     17,
		},
		{fileName: "examples/good/custom.txt",
			reference: &Graph{AntNum: 1,
				Start: &Point{Name: "0", X: 0, Y: 3},
				End:   &Point{Name: "1", X: 1, Y: 3},
				Points: []*Point{
					{Name: "2", X: 2, Y: 8},
				}, Edges: []*Edge{
					{From: &Point{Name: "0"}, To: &Point{Name: "1"}},
					{From: &Point{Name: "0"}, To: &Point{Name: "2"}},
					{From: &Point{Name: "2"}, To: &Point{Name: "1"}},
				},
			},
			solutionLength: 1,
			edgeLength:     3,
		},
	}
	for _, test := range goodTests {
		t.Run("GOOD-"+test.fileName, func(t *testing.T) {
			file, err := os.ReadFile(test.fileName)
			if err != nil {
				t.Fatal(err)
			}
			graph, err := ReadGraph(string(file))
			if err != nil {
				t.Fatalf("expected result, got error %v", err)
			}
			if graph.AntNum != test.reference.AntNum {
				t.Errorf("AntNum is not valid, got %v, expected %v", graph.AntNum, test.reference.AntNum)
			}
			if *graph.Start != *test.reference.Start {
				t.Errorf("Start is not valid, got %v, expected %v", *graph.Start, *test.reference.Start)
			}
			if *graph.End != *test.reference.End {
				t.Errorf("End is not valid, got %v, expected %v", *graph.End, *test.reference.End)
			}

			// TODO: add same test for Edges
			if len(test.reference.Points) != 0 && !reflect.DeepEqual(graph.Points, test.reference.Points) {
				t.Errorf("Points are not valid, got %v, expected %v", graph.Points, test.reference.Points)
			}

			if len(test.reference.Edges) != 0 && !twoEdgeSequencesTheSame(graph.Edges, test.reference.Edges) {
				t.Errorf("Edges are not valid, got %v, expected %v", graph.Edges, test.reference.Edges)
			}

			err = graph.FindPaths()
			if err != nil {
				t.Errorf("can not find paths of valid graph, %v", err)
			}
			solution := graph.Solve()
			solutionLength := len(strings.Split(graph.FormatSolution(solution), "\n"))
			if test.solutionLength != 0 && solutionLength != test.solutionLength {
				t.Errorf("wrong solution length, got %v, expected %v", solutionLength, test.solutionLength)
			}

		})
	}
	badTests, err := filepath.Glob("examples/bad_format/*")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range badTests {
		t.Run("BAD-"+test, func(t *testing.T) {
			file, err := os.ReadFile(test)
			if err != nil {
				t.Fatal(err)
			}
			res, err := ReadGraph(string(file))
			if err == nil {
				t.Fatalf("expected error, got %v", res)
			}
		})
	}
	badPathsTests, err := filepath.Glob("examples/bad_paths/*")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range badPathsTests {
		t.Run("BAD_PATHS-"+test, func(t *testing.T) {
			file, err := os.ReadFile(test)
			if err != nil {
				t.Fatal(err)
			}
			res, err := ReadGraph(string(file))
			if err != nil {
				t.Fatalf("expected valid format, got %v", err)
			}
			err = res.FindPaths()
			if err == nil {
				t.Fatalf("expected error, got %#v", res.Paths)
			}
		})
	}
}
