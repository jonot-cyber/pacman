package main

// Pos A position in a room.
type Pos struct {
	X int // X position from 0-15
	Y int // Y position from 0-15
}

// Wrap If you go over the edge of the room, loop back around
func (p Pos) Wrap() Pos {
	if p.X < 0 {
		p.X = 14
	}
	if p.X > 14 {
		p.X = 0
	}
	if p.Y < 0 {
		p.Y = 14
	}
	if p.Y > 14 {
		p.Y = 0
	}
	return p
}

// Up returns a position 1 unit up
func (p Pos) Up() Pos {
	return Pos{
		p.X,
		p.Y - 1,
	}
}

// Down returns a position 1 unit down
func (p Pos) Down() Pos {
	return Pos{
		p.X,
		p.Y + 1,
	}
}

// Left returns a position 1 unit left
func (p Pos) Left() Pos {
	return Pos{
		p.X - 1,
		p.Y,
	}
}

// Right returns a position 1 unit right
func (p Pos) Right() Pos {
	return Pos{
		p.X + 1,
		p.Y,
	}
}

// Cascade recursively travels the board, recording the shortest path to a certain point.
func (p Pos) Cascade(room *Room, goDirection *map[Pos]Result, direction uint, step uint) {
	if room.SolidAtPos(p) {
		return
	}
	if v, ok := (*goDirection)[p]; ok {
		if step < v.Steps {
			(*goDirection)[p] = Result{Direction: direction, Steps: step}
		} else {
			return // If it is less efficient to go that path, it will be less efficient for all additional steps.
			// Also is end case for recursion
		}
	} else {
		(*goDirection)[p] = Result{direction, step}
	}
	p.Up().Wrap().Cascade(room, goDirection, Down, step+1)
	p.Down().Wrap().Cascade(room, goDirection, Up, step+1)
	p.Left().Wrap().Cascade(room, goDirection, Right, step+1)
	p.Right().Wrap().Cascade(room, goDirection, Left, step+1)
}
