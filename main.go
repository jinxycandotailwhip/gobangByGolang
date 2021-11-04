package main

import (
	"fmt"
	"gobang/negative_max"
	"gobang/variable"
	"log"
	"strconv"
	"strings"
)

// func main() {
// 	variable.List2[variable.Coordinate{1, 1}] = true
// 	variable.List3[variable.Coordinate{1, 1}] = true
// 	variable.List1[variable.Coordinate{3, 3}] = true
// 	variable.List3[variable.Coordinate{3, 3}] = true
// 	variable.List1[variable.Coordinate{4, 4}] = true
// 	variable.List3[variable.Coordinate{4, 4}] = true
// 	variable.List1[variable.Coordinate{5, 5}] = true
// 	variable.List3[variable.Coordinate{5, 5}] = true
// 	variable.List1[variable.Coordinate{6, 6}] = true
// 	variable.List3[variable.Coordinate{6, 6}] = true
// 	negative_max.AI()
// }

func main() {
	grids := make([][]int, 15)
	for i := 0; i < 15; i++ {
		grids[i] = make([]int, 15)
	}
	for {
		fmt.Println("----------- your turn -----------")
		var coor string
		fmt.Println("input coordinate x, y")
		fmt.Scanln(&coor)
		coors := strings.Split(coor, ",")
		x, err := strconv.Atoi(coors[0])
		if err != nil {
			log.Default().Println("x: string to int failed")
		}
		y, err := strconv.Atoi(coors[1])
		if err != nil {
			log.Default().Println("y: string to int failed")
		}
		variable.List1[variable.Coordinate{x, y}] = true
		variable.List3[variable.Coordinate{x, y}] = true
		fmt.Println("(", x, ",", y, ")")
		grids[x][y] = 1
		fmt.Println("**************************************")
		for _, val := range grids {
			fmt.Println(val)
		}
		fmt.Println("**************************************")
		if negative_max.GameWin(variable.List1) {
			fmt.Println("---you win---")
			break
		}
		fmt.Println("----------- AI's turn -----------")
		fmt.Println("list1", variable.List1)
		fmt.Println("list2", variable.List2)
		pos := negative_max.AI()
		fmt.Println(pos)
		fmt.Println("pos: ", pos)
		variable.List2[pos] = true
		variable.List3[pos] = true
		grids[pos[0]][pos[1]] = 2
		fmt.Println("**************************************")
		for _, val := range grids {
			fmt.Println(val)
		}
		fmt.Println("**************************************")
		if negative_max.GameWin(variable.List2) {
			fmt.Println("---ai win---")
			break
		}
	}
}
