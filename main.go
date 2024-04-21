package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type State struct {
	board     [3][3]int
	zeroX     int
	zeroY     int
	cost      int
	heuristic int
	prev      *State
}

var solution = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost+pq[i].heuristic < pq[j].cost+pq[j].heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func readRow(board *[3][3]int, rowNum int) error {
	var input string
	fmt.Scanln(&input)
	numbers, err := parseRow(input)
	if err != nil {
		return err
	}
	copy(board[rowNum][:], numbers[:]) // Convertendo a matriz em uma fatia antes de copiar
	return nil
}

func parseRow(input string) ([3]int, error) {
	var row [3]int
	numbers := strings.Split(input, ",")
	if len(numbers) != 3 {
		return row, fmt.Errorf("invalid input: must contain 3 numbers separated by commas")
	}
	for i := 0; i < 3; i++ {
		fmt.Sscanf(numbers[i], "%d", &row[i])
	}
	return row, nil
}

func HeuristicSimple(state *State) int {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state.board[i][j] != 0 && state.board[i][j] != solution[i][j] {
				count++
			}
		}
	}
	return count
}

func HeuristicManhattan(state *State) int {
	targetPos := map[int][2]int{
		1: {0, 0}, 2: {0, 1}, 3: {0, 2},
		4: {1, 0}, 5: {1, 1}, 6: {1, 2},
		7: {2, 0}, 8: {2, 1}, 0: {2, 2},
	}
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			value := state.board[i][j]
			if value != 0 {
				target := targetPos[value]
				sum += abs(target[0]-i) + abs(target[1]-j)
			}
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindZero(board [3][3]int) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // this should never happen if input is correct
}

func Expand(current *State) []*State {
	directions := []struct{ dx, dy int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var states []*State
	fmt.Println("Expanding states...")
	for _, d := range directions {
		nx, ny := current.zeroX+d.dx, current.zeroY+d.dy
		if nx >= 0 && nx < 3 && ny >= 0 && ny < 3 {
			newBoard := current.board
			newBoard[current.zeroX][current.zeroY], newBoard[nx][ny] = newBoard[nx][ny], newBoard[current.zeroX][current.zeroY]
			newState := &State{
				board:     newBoard,
				zeroX:     nx,
				zeroY:     ny,
				cost:      current.cost + 1,
				heuristic: 0, // Will be set after choosing the heuristic
				prev:      current,
			}
			states = append(states, newState)
		}
	}
	return states
}

func main() {
	var board [3][3]int
	fmt.Println("Enter the each row separatedly. 0 represents the empty space.")
	fmt.Println("Enter the first row of the puzzle board, separate by commas (e.g., 1,2,3)")
	if err := readRow(&board, 0); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter the second row of the puzzle board:")
	if err := readRow(&board, 1); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter the third row of the puzzle board:")
	if err := readRow(&board, 2); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Agora você tem o tabuleiro montado na matriz 'board'
	fmt.Println("Board:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}

	if !Solvable(board, solution) {
		log.Fatalf("Sem solução")
	}

	zeroX, zeroY := FindZero(board)
	fmt.Printf("Zero found at (%d, %d)\n", zeroX, zeroY)

	startState := &State{
		board:     board,
		zeroX:     zeroX,
		zeroY:     zeroY,
		cost:      0,
		heuristic: 0,
		prev:      nil,
	}

	fmt.Println("Choose the search algorithm: 1 for Uniform Cost, 2 for A* Simple, 3 for A* Precise")
	var choice int
	fmt.Scan(&choice)

	var heuristicFunc func(*State) int
	switch choice {
	case 1:
		heuristicFunc = func(s *State) int { return 0 } // Uniform Cost Search
	case 2:
		heuristicFunc = HeuristicSimple
	case 3:
		heuristicFunc = HeuristicManhattan
	default:
		fmt.Println("Invalid choice.")
		os.Exit(1)
	}

	startState.heuristic = heuristicFunc(startState)
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, startState)

	nodesVisited := 0
	startTime := time.Now()

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)
		nodesVisited++

		fmt.Printf("Visiting node with heuristic %d and cost %d\n", current.heuristic, current.cost)

		if current.heuristic == 0 { // Check if this is the goal state
			duration := time.Since(startTime)
			fmt.Println("Solution path:")
			printSolution(current)
			fmt.Printf("Nodes visited: %d\n", nodesVisited)
			fmt.Printf("Path length: %d\n", current.cost)
			fmt.Printf("Execution time: %fs\n", duration.Seconds())
			return
		}

		for _, nextState := range Expand(current) {
			nextState.heuristic = heuristicFunc(nextState)
			heap.Push(&pq, nextState)
		}
	}

	fmt.Println("No solution found")
}

func printSolution(state *State) {
	if state.prev != nil {
		printSolution(state.prev)
	}
	fmt.Println(state.board)
}

// Solvable receives two boards and returns if the game is solvable
func Solvable(board [3][3]int, goal [3][3]int) bool {
	var invGoal int
	var boardArray, goalArray [9]int

	counter := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardArray[counter] = board[i][j]
			goalArray[counter] = goal[i][j]
			counter++
		}
	}

	fmt.Println(boardArray)

	for i := range goalArray {
		for j := i + 1; j < len(goal); j++ {
			if goalArray[i] > goalArray[j] && goalArray[j] != 0 {
				invGoal++
			}
		}
	}
	invBoard := 0
	for i := range boardArray {
		for j := i + 1; j < len(boardArray); j++ {
			if boardArray[i] > boardArray[j] && boardArray[j] != 0 {
				invBoard++
			}
		}
	}
	return invGoal%2 == invBoard%2
}
