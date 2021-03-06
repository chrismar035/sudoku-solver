package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chrismar035/sudoku-solver"
)

func main() {
	file, err := os.Open("puzzles.txt")
	if err != nil {
		fmt.Println("Error opening puzzles", err)
	}
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if len(line) != 81 {
			fmt.Println("Puzzle improperly sized on line", line, err)
		}

		var grid solver.Grid
		for i, char := range strings.Split(line, "") {
			grid[i], err = strconv.Atoi(char)
			if err != nil {
				fmt.Println("Invalid puzzle line", lineNumber, "char", i)
			}
		}

		puzzle := solver.Puzzle{Initial: grid}
		multiSolver := solver.NewMultiBacktrackingSolver()
		solutions, _ := multiSolver.Solve(grid)

		fmt.Println("Found", len(solutions), "solution(s)")
		fmt.Print("Initial\n", puzzle.Initial, "\n\nSolution(s)\n")
		for _, g := range solutions {
			fmt.Print(g)
			fmt.Println("\n-----------\n")
		}

		singleSolver := solver.NewSolver()
		puzzle.Solution, _ = singleSolver.Solve(grid)
		fmt.Println(puzzle)

	}
}
