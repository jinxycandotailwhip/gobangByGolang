package server

import "gobang/variable"

func Grid2index(grid variable.Coordinate) int {
	return grid[0] + 15*grid[1]
}

func Index2grid(index int) variable.Coordinate {
	return variable.Coordinate{
		index % 15,
		index / 15,
	}
}
