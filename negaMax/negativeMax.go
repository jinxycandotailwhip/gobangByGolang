package negative_max

import (
	"fmt"
	"gobang/variable"
	"math"
)

func AI() [2]int {
	_ = negaMax(true, variable.DEPTH, math.MinInt, math.MaxInt)
	fmt.Println("本次共剪枝次数: ", variable.CutCount)
	fmt.Println("本次共搜索次数: ", variable.SearchCount)
	return variable.NextPoint
}

func negaMax(is_ai bool, depth, alpha, beta int) int {
	// 判断游戏是否结束 || 搜索深度是否达到
	if gameWin(variable.List1) || gameWin(variable.List2) || depth == 0 {

	}
}

func gameWin(list []int) {
	for m:=0; i<variable.COLUMN; i++ {
		for n:=0; j<variable.ROW; j++ {
			if n < variable.ROW - 4 && 
		}
	}
}
