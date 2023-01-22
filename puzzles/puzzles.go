package puzzles

type Puzzle struct {
	Name     string
	Part     string
	Solution int64
	Input    []string
}

func (p *Puzzle) resolve() {}