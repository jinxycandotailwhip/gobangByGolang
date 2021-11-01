package variable

type coordinate [2]int

var (
	List1 = map[coordinate]bool{} // AI
	List2 = map[coordinate]bool{} // human
	List3 = map[coordinate]bool{} // all

	ListAll   = []int{}      // 整个棋盘的点
	NextPoint = [2]int{0, 0} // AI下一步最应该下的点

	CutCount    = 0 // 统计剪枝次数
	SearchCount = 0 // 统计搜索次数
)
