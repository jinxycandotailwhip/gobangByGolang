package variable

type Coordinate [2]int

var (
	List1 = map[Coordinate]bool{} // AI
	List2 = map[Coordinate]bool{} // human
	List3 = map[Coordinate]bool{} // all

	ListAll   = []int{}      // 整个棋盘的点
	NextPoint = [2]int{0, 0} // AI下一步最应该下的点

	CutCount    = 0 // 统计剪枝次数
	SearchCount = 0 // 统计搜索次数
)
