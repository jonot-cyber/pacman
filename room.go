package main

type StupidType map[Pos]*map[Pos]Result

type Room struct {
	Solid     [15][15]bool
	Direction map[Pos]Result
}

func (r Room) SolidAtPos(p Pos) bool {
	return r.Solid[p.Y][p.X]
}

func (r Room) GenerateMapWithGoal(g Pos) *map[Pos]Result {
	var ret = make(map[Pos]Result)
	g.Cascade(&r, &ret, Up, 0)
	return &ret
}

func (r Room) GenerateMap() StupidType {
	var stupidType = make(map[Pos]*map[Pos]Result)
	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
			pos := Pos{x, y}
			if r.SolidAtPos(pos) {
				continue
			}
			stupidType[pos] = r.GenerateMapWithGoal(pos)
		}
	}
	return stupidType
}
