package rcss

type Direction float64

type SightWidth string

type Power int

type Moment int

type NeckMoment int

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
	Dash(p Power) error
	Kick(p Power, d Direction) error
	Move(x, y int) error
	Say(msg string) error
	Turn(m Moment) error
	TurnNeck(n NeckMoment) error

	Score() error

	See() error
	SenseBody() error
}
