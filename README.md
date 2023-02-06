# task-lem-in
grit:lab Åland Islands 2022

---
Authors: [@maximihajlov](https://github.com/maximihajlov), [@healingdrawing](https://github.com/healingdrawing), [@nattikim](https://github.com/nattikim)

Solved during studying in Gritlab coding school on Åland, December 2022

---

## [Task description and audit questions](https://github.com/01-edu/public/tree/master/subjects/lem-in)

## Usage

### Run `go run lem-in [FILE]`

To read graph configuration from `[FILE]` and solve it. Solution will be printed to Stdout

### Example: `go run lem-in ./examples/example00.txt`

### Visualizer

### Run `go run lem-in/visualizatizer [OUTPUT_FILE.html]`

Write solution to Stdin to visualize it

Solution will be saved as HTML page with an SVG image to `[OUTPUT_FILE]`(or default `./solution.html`)

### Example: `go run lem-in ./examples/example00.txt | go run lem-in/visualizer`

To solve and visualize `example00.txt`

### Testing: `go test lem-in/solver/test -cover -coverpkg=./...`
