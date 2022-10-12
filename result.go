package main

const (
	Up uint = iota
	Down
	Left
	Right
)

// Result of pathfinding
type Result struct {
	Direction uint // Direction
	Steps     uint // Number of steps to target.
}
