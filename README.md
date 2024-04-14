# 8-Puzzle Solver in Go

This project implements an 8-puzzle solver in Go, providing three different search algorithms to solve the puzzle: Uniform Cost Search, A* with a simple heuristic, and A* with the Manhattan distance heuristic. The program reads the puzzle state from the user and outputs the solution path, the number of nodes visited, the path length, and the execution time.

## Prerequisites

To run this program, you will need:

- Go (version 1.15 or higher)

You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

## Installation

To install this program, clone the repository to your local machine:

```bash
git clone https://github.com/rafaelescrich/8-puzzle-solver.git
cd 8-puzzle-solver
```

## Usage

To run the program, navigate to the directory containing the `main.go` file and run:

```bash
go run main.go
```

You will be prompted to enter the 8-puzzle board configuration as a series of numbers from 0-8, where 0 represents the empty space. For example:

```bash
Enter the 3x3 puzzle board (0 represents the empty space):
1 2 3
4 5 6
7 8 0
```

Next, choose the search algorithm:

```bash
Choose the search algorithm: 1 for Uniform Cost, 2 for A* Simple, 3 for A* Precise
```

The program will process the input and display the solution path, if one is found, along with performance statistics.

## Contributing

To contribute, please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md)  file for details.
