package test

import (
	"gobang/negative_max"
	"gobang/variable"
	"testing"
)

func TestHasNeighbour(t *testing.T) {
	variable.List3[variable.Coordinate{4, 4}] = true
	if !negative_max.HasNeighbour(variable.Coordinate{3, 3}) {
		t.Error("left-top neighbour not detected")
	} else {
		t.Log("point left-top passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{3, 4}) {
		t.Error("left neighbour not detected")
	} else {
		t.Log("point left passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{3, 5}) {
		t.Error("left-bottom neighbour not detected")
	} else {
		t.Log("point left-bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{4, 5}) {
		t.Error("bottom neighbour not detected")
	} else {
		t.Log("point bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 5}) {
		t.Error("right-bottom neighbour not detected")
	} else {
		t.Log("point right-bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 4}) {
		t.Error("right neighbour not detected")
	} else {
		t.Log("point right passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 3}) {
		t.Error("right-top neighbour not detected")
	} else {
		t.Log("point right-top passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{4, 3}) {
		t.Error("top neighbour not detected")
	} else {
		t.Log("point top passed")
	}
	if negative_max.HasNeighbour(variable.Coordinate{6, 6}) {
		t.Error("false neightbour detected")
	} else {
		t.Log("point not neighbour passed")
	}
}
