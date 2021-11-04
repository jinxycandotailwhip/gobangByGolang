package test

import (
	"gobang/negative_max"
	"gobang/variable"
	"testing"
)

var coors1 = map[variable.Coordinate]bool{
	{0, 0}: true,
	{1, 1}: true,
	{2, 2}: true,
	{3, 3}: true,
	{4, 4}: true,
}
var coors2 = map[variable.Coordinate]bool{
	{4, 2}: true,
	{4, 3}: true,
	{4, 4}: true,
	{4, 5}: true,
	{4, 6}: true,
}
var coors3 = map[variable.Coordinate]bool{
	{8, 8}: true,
	{7, 7}: true,
	{6, 6}: true,
	{5, 5}: true,
	{4, 4}: true,
}

var coors4 = map[variable.Coordinate]bool{
	{8, 4}: true,
	{7, 7}: true,
	{6, 6}: true,
	{5, 5}: true,
	{4, 4}: true,
}

func TestGameWin(t *testing.T) {
	if !negative_max.GameWin(coors1) {
		t.Error("coor1 expected gamewin is true")
	} else {
		t.Log("coor1 passed")
	}
	if !negative_max.GameWin(coors2) {
		t.Error("coor2 expected gamewin is true")
	} else {
		t.Log("coor2 passed")
	}
	if !negative_max.GameWin(coors3) {
		t.Error("coor3 expected gamewin is true")
	} else {
		t.Log("coor3 passed")
	}
	if negative_max.GameWin(coors4) {
		t.Error("coor4 expected gamewin is false")
	} else {
		t.Log("coor4 passed")
	}
}
