package flatdom

// Grid is a unit on the land. It contains information of who holds
// this piece of land and what is its influence.
// When influence drops to 0, this grid is declared as a no-man land.
type Grid struct {
	influence float64
	holder    int
}
