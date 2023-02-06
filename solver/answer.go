package solver

import (
	"strconv"
	"strings"
)

// FormatSolution returns an instructions for each ant in required format
//
// Check task description for more info
func (g *Graph) FormatSolution(solution []*Path) string {
	var lines = make(map[int][]string)

	for ant := 0; ant < g.AntNum; ant++ {
		antPath := solution[ant%len(solution)]
		pathStartLine := ant/len(solution) + 1

		if g.AntNum-len(solution) < ant {
			// fix for example02.txt where it's better to wait when better path will be available
			// (and start ant later as a antNew=ant+1)
			for antNew := ant + 1; antNew < ant+len(solution); antNew++ {
				newAntPath := solution[antNew%len(solution)]
				newPathStartLine := antNew/len(solution) + 1

				if len(*newAntPath)+newPathStartLine < len(*antPath)+newPathStartLine {
					antPath = newAntPath
					pathStartLine = newPathStartLine
				}
			}
		}

		for i, point := range (*antPath)[1:] { // first is always a start
			lines[pathStartLine+i] = append(lines[pathStartLine+i], "L"+strconv.Itoa(ant+1)+"-"+point.Name)
		}
	}

	var linesSlice []string
	for i := 1; i <= len(lines); i++ {
		linesSlice = append(linesSlice, strings.Join(lines[i], " "))
	}
	return strings.Join(linesSlice, "\n")
}
