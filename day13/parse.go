package main

import (
	"strconv"
	"strings"

	"advent2024/util"
)

func processPuzzles() []puzzle {
	var puzzles []puzzle

	currentPuzzle := &puzzle{}

	for line := range util.Data(13) {
		if currentPuzzle == nil {
			currentPuzzle = &puzzle{}
		}

		if strings.HasPrefix(line, "Button") {
			parts := strings.Split(line, ": ")

			variables := strings.Split(parts[1], ", ")
			var currentChunk chunk
			for _, variable := range variables {
				getChunk(&currentChunk, variable, "+")
			}

			if strings.Contains(parts[0], "A") {
				currentPuzzle.A = currentChunk
			}
			if strings.Contains(parts[0], "B") {
				currentPuzzle.B = currentChunk
			}
		}

		if strings.HasPrefix(line, "Prize: ") {
			line = strings.TrimPrefix(line, "Prize: ")
			var currentChunk chunk
			variables := strings.Split(line, ", ")
			for _, variable := range variables {
				getChunk(&currentChunk, variable, "=")
			}
			currentPuzzle.Prize = currentChunk
		}

		if line == "" {
			puzzles = append(puzzles, *currentPuzzle)
			currentPuzzle = nil
		}
	}

	if currentPuzzle != nil {
		puzzles = append(puzzles, *currentPuzzle)
	}

	return puzzles
}

func getChunk(c *chunk, variable string, sep string) {
	parts := strings.Split(variable, sep)

	if parts[0] == "X" {
		c.X = util.Must(strconv.Atoi(parts[1]))
	}
	if parts[0] == "Y" {
		c.Y = util.Must(strconv.Atoi(parts[1]))
	}
}
