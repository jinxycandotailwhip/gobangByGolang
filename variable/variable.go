package variable

type Coordinate [2]int
type SessionInfo struct {
	List1       map[Coordinate]bool
	List2       map[Coordinate]bool
	List3       map[Coordinate]bool
	ListAll     []int
	NextPoint   [2]int
	CutCount    int
	SearchCount int
	Finish      bool
}

var (
	SessionMap = map[string]*SessionInfo{}
)
