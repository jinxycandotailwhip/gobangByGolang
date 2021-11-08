package negative_max

import (
	"fmt"
	"gobang/variable"

	"go.uber.org/zap"
)

type scoreShape struct {
	score  int
	points [][2]int
	direct [2]int
}

func AI(session *variable.SessionInfo) [2]int {
	fmt.Println(len(variable.SessionMap))
	_ = negaMax(true, variable.DEPTH, -9999999, 9999999, session)
	fmt.Println("本次共剪枝次数: ", session.CutCount)
	fmt.Println("本次共搜索次数: ", session.SearchCount)
	return session.NextPoint
}

func negaMax(isAI bool, depth, alpha, beta int, session *variable.SessionInfo) int {
	// 判断游戏是否结束 || 搜索深度是否达到
	if GameWin(session.List1) || GameWin(session.List2) || depth == 0 {
		return evaluation(isAI, session)
	}
	// blankList是棋面上还没有下的点
	blankList := make(map[variable.Coordinate]bool)
	for m := 0; m < 15; m++ {
		for n := 0; n < 15; n++ {
			if !session.List3[variable.Coordinate{m, n}] {
				blankList[variable.Coordinate{m, n}] = true
			}
		}
	}
	for nextStep := range blankList {
		session.SearchCount++
		// 如果没有相邻点就不用考虑
		if !HasNeighbour(nextStep, session) {
			continue
		}
		if isAI {
			session.List1[nextStep] = true
		} else {
			session.List2[nextStep] = true
		}
		session.List3[nextStep] = true

		value := -negaMax(!isAI, depth-1, -beta, -alpha, session)
		if isAI {
			delete(session.List1, nextStep)
		} else {
			delete(session.List2, nextStep)
		}
		delete(session.List3, nextStep)

		if value > alpha {
			// fmt.Println(value, "alpha:", alpha, "beta:", beta)
			if depth == variable.DEPTH {
				session.NextPoint = nextStep
			}
			// alpha beta 剪枝点
			if value >= beta {
				session.CutCount++
				return beta
			}
			alpha = value
		}
	}
	return alpha
}

func GameWin(list map[variable.Coordinate]bool) bool {
	for m := 0; m < variable.COLUMN; m++ {
		for n := 0; n < variable.ROW; n++ {
			if n < variable.ROW-4 &&
				list[[2]int{m, n}] &&
				list[[2]int{m, n + 1}] &&
				list[[2]int{m, n + 2}] &&
				list[[2]int{m, n + 3}] &&
				list[[2]int{m, n + 4}] {
				return true
			} else if m < variable.ROW-4 &&
				list[[2]int{m, n}] &&
				list[[2]int{m + 1, n}] &&
				list[[2]int{m + 2, n}] &&
				list[[2]int{m + 3, n}] &&
				list[[2]int{m + 4, n}] {
				return true
			} else if m < variable.ROW-4 && n < variable.ROW-4 &&
				list[[2]int{m, n}] &&
				list[[2]int{m + 1, n + 1}] &&
				list[[2]int{m + 2, n + 2}] &&
				list[[2]int{m + 3, n + 3}] &&
				list[[2]int{m + 4, n + 4}] {
				return true
			} else if m < variable.ROW-4 && n > 3 &&
				list[[2]int{m, n}] &&
				list[[2]int{m + 1, n - 1}] &&
				list[[2]int{m + 2, n - 2}] &&
				list[[2]int{m + 3, n - 3}] &&
				list[[2]int{m + 4, n - 4}] {
				return true
			}
		}
	}
	return false
}

func evaluation(isAI bool, session *variable.SessionInfo) int {
	var myList, enemyList map[variable.Coordinate]bool
	if isAI {
		myList = session.List1
		enemyList = session.List2
	} else {
		myList = session.List2
		enemyList = session.List1
	}

	// 计算自己的得分
	scoreAllArr := make(map[*scoreShape]bool)
	myScore := 0
	for coor := range myList {
		m, n := coor[0], coor[1]
		myScore += calScore(m, n, 0, 1, enemyList, myList, scoreAllArr)
		myScore += calScore(m, n, 1, 0, enemyList, myList, scoreAllArr)
		myScore += calScore(m, n, 1, 1, enemyList, myList, scoreAllArr)
		myScore += calScore(m, n, -1, 1, enemyList, myList, scoreAllArr)
	}

	// 计算敌方的得分
	scoreAllArrEnemy := make(map[*scoreShape]bool)
	enemyScore := 0
	for coor := range enemyList {
		m, n := coor[0], coor[1]
		enemyScore += calScore(m, n, 0, 1, myList, enemyList, scoreAllArrEnemy)
		enemyScore += calScore(m, n, 1, 0, myList, enemyList, scoreAllArrEnemy)
		enemyScore += calScore(m, n, 1, 1, myList, enemyList, scoreAllArrEnemy)
		enemyScore += calScore(m, n, -1, 1, myList, enemyList, scoreAllArrEnemy)
	}

	// 我方分数减去敌方分数*比例
	return int(float32(myScore) - float32(enemyScore*variable.RATIO)*0.1)
}

// 每个方向上的分值计算，一个方向上只有可能有一种得分，calScore在被调用的时候会调用四次，也就是四个方向，不用方向的得分会叠加
func calScore(m, n, xDirect, yDirect int, enemyList, myList map[variable.Coordinate]bool, scoreAllArr map[*scoreShape]bool) int {
	var maxScoreShape scoreShape
	// ss stand for score shape
	for ss := range scoreAllArr {
		for _, coor := range ss.points {
			// 判断(m, n)是否是scoreAllArr里的点，也就是判断(m, n)在当前方向是否已经被计算过分数，如果已经计算过，不重复计算，返回0值
			if m == coor[0] && n == coor[1] && xDirect == ss.direct[0] && yDirect == ss.direct[1] {
				return 0
			}
		}
	}
	// 在落子点的左右方向上循环查找得分形状
	for offset := -5; offset < 1; offset++ {
		pos := []int{}
		for i := 0; i < 6; i++ {
			if enemyList[variable.Coordinate{m + (i+offset)*xDirect, n + (i+offset)*yDirect}] {
				pos = append(pos, 2)
			} else if myList[variable.Coordinate{m + (i+offset)*xDirect, n + (i+offset)*yDirect}] {
				pos = append(pos, 1)
			} else {
				pos = append(pos, 0)
			}
		}
		tmpShape5 := [5]int{pos[0], pos[1], pos[2], pos[3], pos[4]}
		tmpShape6 := [6]int{pos[0], pos[1], pos[2], pos[3], pos[4], pos[5]}
		if score, ok := variable.ShapeScoreList5[tmpShape5]; ok {
			// 当前方向上只取最大得分，如果当前score比之前计算的得分大，那么覆盖之前的得分（三连取代二连）
			if score > maxScoreShape.score {
				if tmpShape5 == [5]int{1, 1, 1, 1, 1} {
					zap.L().Info("win!!!!!!!!!!!")
				}
				maxScoreShape.score = score
				maxScoreShape.points = [][2]int{
					{m + (0+offset)*xDirect, n + (0+offset)*yDirect},
					{m + (1+offset)*xDirect, n + (1+offset)*yDirect},
					{m + (2+offset)*xDirect, n + (2+offset)*yDirect},
					{m + (3+offset)*xDirect, n + (3+offset)*yDirect},
					{m + (4+offset)*xDirect, n + (4+offset)*yDirect},
				}
				maxScoreShape.direct = [2]int{xDirect, yDirect}
			}
		}
		if score, ok := variable.ShapeScoreList6[tmpShape6]; ok {
			// 当前方向上只取最大得分，如果当前score比之前计算的得分大，那么覆盖之前的得分（三连取代二连）
			if score > maxScoreShape.score {
				maxScoreShape.score = score
				maxScoreShape.points = [][2]int{
					{m + (0+offset)*xDirect, n + (0+offset)*yDirect},
					{m + (1+offset)*xDirect, n + (1+offset)*yDirect},
					{m + (2+offset)*xDirect, n + (2+offset)*yDirect},
					{m + (3+offset)*xDirect, n + (3+offset)*yDirect},
					{m + (4+offset)*xDirect, n + (4+offset)*yDirect},
					// FIXME: 这里是不是应该还有个5+offset，因为这里的评分是用6个格子评的
				}
				maxScoreShape.direct = [2]int{xDirect, yDirect}
			}
		}
	}
	return maxScoreShape.score
}

func HasNeighbour(coor variable.Coordinate, session *variable.SessionInfo) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// i==0 && j==0是这个点本身
			if i == 0 && j == 0 {
				continue
			}
			cur := variable.Coordinate{
				coor[0] + i,
				coor[1] + j,
			}
			if session.List3[cur] {
				return true
			}
		}
	}
	return false
}
