package main

import (
	"fmt"
	"testing"
)

func TestHeuristicUniform(t *testing.T) {
	// Test cases don't matter since the output is always 0
	state1 := &State{board: [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}}
	state2 := &State{board: [3][3]int{{8, 6, 7}, {2, 5, 4}, {3, 0, 1}}}

	if HeuristicUniform(state1) != 0 || HeuristicUniform(state2) != 0 {
		t.Error("HeuristicUniform should always return 0")
	}
}

func TestHeuristicSimple(t *testing.T) {
	board := [3][3]int{{8, 1, 3}, {0, 4, 2}, {7, 6, 5}} // Input board
	expectedCount := 6                                  // I think the problem might be how I calculate this
	result := HeuristicSimple(&State{board: board})

	if result != expectedCount {
		t.Errorf("HeuristicSimple failed. Expected: %d, Got: %d", expectedCount, result)
	}
}

func TestHeuristicManhattan(t *testing.T) {
	state := &State{
		board: [3][3]int{{8, 1, 3}, {4, 0, 2}, {7, 6, 5}},
	}
	expectedDistance := 10 // Calculate this manually based on the state
	distance := HeuristicManhattan(state)

	if distance != expectedDistance {
		t.Errorf("HeuristicManhattan returned %d, expected %d", distance, expectedDistance)
	}
}

// TestSolvable checks if the function correctly identifies solvable puzzles.
func TestSolvable(t *testing.T) {
	board := [3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}} // Known solvable position
	if !Solvable(board, solution) {
		t.Errorf("Solvable returned false; want true for a known solvable configuration")
	}
}

func TestExpand(t *testing.T) {
	state := &State{board: [3][3]int{{1, 2, 3}, {4, 5, 0}, {7, 8, 6}}, zeroX: 2, zeroY: 1, cost: 0}

	fmt.Println("Original Board:") // Print for clarity
	for _, row := range state.board {
		fmt.Println(row)
	}

	newStates := Expand(state)
	expectedNumMoves := 3

	if len(newStates) != expectedNumMoves {
		t.Errorf("Expand returned %d possible states; want %d", len(newStates), expectedNumMoves)
	}
}
