package rcss

type Side byte

const (
	LeftSide Side = 'l'
	RightSide Side = 'r'
)

type UniformNumber uint8

type PlayMode string

const (
	BeforeKickOff PlayMode = "before_kick_off"
	PlayOn PlayMode = "play_on"
)

// Input Driver
type Team interface {
	Name() string

	Init(match Match, side Side, unum UniformNumber, mode PlayMode)
	ServerParam()
	PlayerParam()
	PlayerType()

	See()
	Hear()
	SenseBody()

	Score()
}
