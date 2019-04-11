package rcss

type Direction float64

type SightWidth string

const (
	NarrowSight SightWidth = "narrow"
	NormalSight SightWidth = "normal"
	WideSight   SightWidth = "wide"
)

type SightQuality string

const (
	LowVision  SightQuality = "low"
	HighVision SightQuality = "high"
)

// Output Driver
type Match interface {
	Join(team Team) error
	Reconnect(team Team, unum UniformNumber) error
	Bye() error

	Catch(dir Direction) error
	ChangeView(w SightWidth, q SightQuality) error
	Dash() error
	Kick() error
	Move(x, y int) error
	Say() error
	Turn() error
	TurnNeck() error

	Score() error

	See() error
	SenseBody() error
}
