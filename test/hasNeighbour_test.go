/**
this file is to test function HasNeighbour
**/
package test

import (
	"gobang/negative_max"
	"gobang/variable"
	"testing"
)

func TestHasNeighbour(t *testing.T) {
	session := variable.SessionInfo{
		List1:       make(map[variable.Coordinate]bool),
		List2:       make(map[variable.Coordinate]bool),
		List3:       make(map[variable.Coordinate]bool),
		ListAll:     []int{},
		NextPoint:   [2]int{0, 0},
		CutCount:    0,
		SearchCount: 0,
	}
	session.List3[variable.Coordinate{4, 4}] = true
	if !negative_max.HasNeighbour(variable.Coordinate{3, 3}, &session) {
		t.Error("left-top neighbour not detected")
	} else {
		t.Log("point left-top passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{3, 4}, &session) {
		t.Error("left neighbour not detected")
	} else {
		t.Log("point left passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{3, 5}, &session) {
		t.Error("left-bottom neighbour not detected")
	} else {
		t.Log("point left-bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{4, 5}, &session) {
		t.Error("bottom neighbour not detected")
	} else {
		t.Log("point bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 5}, &session) {
		t.Error("right-bottom neighbour not detected")
	} else {
		t.Log("point right-bottom passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 4}, &session) {
		t.Error("right neighbour not detected")
	} else {
		t.Log("point right passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{5, 3}, &session) {
		t.Error("right-top neighbour not detected")
	} else {
		t.Log("point right-top passed")
	}
	if !negative_max.HasNeighbour(variable.Coordinate{4, 3}, &session) {
		t.Error("top neighbour not detected")
	} else {
		t.Log("point top passed")
	}
	if negative_max.HasNeighbour(variable.Coordinate{6, 6}, &session) {
		t.Error("false neightbour detected")
	} else {
		t.Log("point not neighbour passed")
	}
}
