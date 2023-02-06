package solver

import (
	"fmt"
	"math"
	"sort"
)

// FindPaths finds all Paths from Start to End using Depth-first search algorithm
//
// it returns an error if there are no paths
func (g *Graph) FindPaths() error {
	g.ParseChildren()

	isVisited := make(map[*Point]bool)
	prev := make(map[*Point]*Point)

	// DFS (Deep-first search algorithm) traverses graph and finds all paths from point to g.End
	var DFS func(point *Point)
	DFS = func(point *Point) {
		if point == g.End {
			path := Path{g.End}
			for point := prev[point]; point != nil; point = prev[point] {
				path = append(path, point)
			}

			for i := 0; i < len(path)/2; i++ {
				// reverse path
				path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
			}

			g.Paths = append(g.Paths, &path)
			return
		}
		isVisited[point] = true
		for _, child := range g.Children[point] {
			if !isVisited[child] {
				prev[child] = point
				DFS(child)
			}
		}
		isVisited[point] = false
	}

	DFS(g.Start)

	sort.Slice(g.Paths, func(i, j int) bool {
		return len(*g.Paths[i]) < len(*g.Paths[j])
	})

	if len(g.Paths) == 0 {
		return fmt.Errorf("there are no paths from start to end")
	}
	return nil
}

// Solve finds the best solution for given Graph with found Paths
func (g *Graph) Solve() []*Path {
	var groups [][]*Path

	// recursive function that creates a group and tries to add paths to it
	// if it's possible to add path to group, it adds it and calls itself to add next path
	// if it's impossible/useless to add any other paths, it saves group to groups
	var Find func(paths []*Path, group []*Path)

	Find = func(paths []*Path, group []*Path) {
		if len(group) == g.AntNum {
			// more paths won't make solution better because all the ants already have their own parallel paths
			groups = append(groups, group)
			return
		}
		for i, path := range paths {
			for _, takenPath := range group {
				for _, point := range *path {
					// check if all points from path are not taken
					// first and last are always start and end
					for _, takenPoint := range (*takenPath)[1 : len(*takenPath)-1] {
						if point == takenPoint {
							goto NextPoint
						}
					}
				}
			}
			// remove path from paths and add it to group, call itself
			Find(append(paths[:i], paths[i:]...), append(group, path))
		NextPoint:
		}

		if len(group) != 0 {
			groups = append(groups, group)
		}
	}

	// if there's a direct path from Start to End, add it here and don't try to add it in future
	// if not, program will send all the ants to the best direct Start-End route ;)
	if len(*g.Paths[0]) == 2 {
		Find(g.Paths[1:], []*Path{g.Paths[0]})
	} else {
		Find(g.Paths, []*Path{})
	}

	var solution []*Path
	solutionTurns := math.MaxInt
	for _, group := range groups {
		// Finding solution with the smallest amount of turns
		// Example: we have three parallel paths with length A, B and C
		// -- A --
		// -- B --
		// -- C --
		// Let's say that we need `X` turns to move all the AntNum ants from Start to End
		// Max amount of ants, going through path for X turns is: X/length
		// So we have the equation X/A + X/B + X/C = AntNum
		// or X*B*C + X*A*C + X*A*B = AntNum*A*B*C
		// or X * (B*C + A*C + A*B) = AntNum*A*B*C

		bandwidth := 1
		for _, path := range group {
			bandwidth *= len(*path) // A * B * C * ...
		}

		multiSum := 0
		for _, path := range group {
			multiSum += bandwidth / len(*path) // A*B*C/A = B*C -> B*C + A*C + A*B ...
		}

		turns := g.AntNum * bandwidth / multiSum

		if turns < solutionTurns {
			solutionTurns = turns
			solution = group
		}
	}

	return solution
}
