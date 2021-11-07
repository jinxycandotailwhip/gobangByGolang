package main

import "gobang/server"

func main() {
	server.StartServer()
}

// func main() {
// 	grids := make([][]int, 15)
// 	for i := 0; i < 15; i++ {
// 		grids[i] = make([]int, 15)
// 	}
// 	for {
// 		fmt.Println("----------- your turn -----------")
// 		var coor string
// 		fmt.Println("input coordinate x, y")
// 		fmt.Scanln(&coor)
// 		coors := strings.Split(coor, ",")
// 		x, err := strconv.Atoi(coors[0])
// 		if err != nil {
// 			log.Default().Println("x: string to int failed")
// 		}
// 		y, err := strconv.Atoi(coors[1])
// 		if err != nil {
// 			log.Default().Println("y: string to int failed")
// 		}
// 		variable.List2[variable.Coordinate{x, y}] = true
// 		variable.List3[variable.Coordinate{x, y}] = true
// 		fmt.Println("(", x, ",", y, ")")
// 		grids[x][y] = 1
// 		fmt.Println("**************************************")
// 		for _, val := range grids {
// 			fmt.Println(val)
// 		}
// 		fmt.Println("**************************************")
// 		if negative_max.GameWin(variable.List2) {
// 			fmt.Println("---you win---")
// 			break
// 		}
// 		fmt.Println("----------- AI's turn -----------")
// 		pos := negative_max.AI()
// 		fmt.Println(pos)
// 		fmt.Println("pos: ", pos)
// 		variable.List1[pos] = true
// 		variable.List3[pos] = true
// 		grids[pos[0]][pos[1]] = 2
// 		fmt.Println("**************************************")
// 		for _, val := range grids {
// 			fmt.Println(val)
// 		}
// 		fmt.Println("**************************************")
// 		if negative_max.GameWin(variable.List1) {
// 			fmt.Println("---ai win---")
// 			break
// 		}
// 	}
// }
