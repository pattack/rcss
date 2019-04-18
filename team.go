package rcss

import (
	"fmt"
)

type Side rune

const (
	LeftSide  Side = 'l'
	RightSide Side = 'r'
)

type UniformNumber uint8

type PlayMode string

const (
	Null            PlayMode = "null"
	BeforeKickOff   PlayMode = "before_kick_off"
	TimeOver        PlayMode = "time_over"
	PlayOn          PlayMode = "play_on"
	KickOffLeft     PlayMode = "kick_off_left"
	KickOffRight    PlayMode = "kick_off_right"
	KickInLeft      PlayMode = "kick_in_left"
	KickInRight     PlayMode = "kick_in_right"
	FreeKickLeft    PlayMode = "free_kick_left"
	FreeKickRight   PlayMode = "free_kick_right"
	CornerKickLeft  PlayMode = "corner_kick_left"
	CornerKickRight PlayMode = "corner_kick_right"
	GoalKickLeft    PlayMode = "goal_kick_left"
	GoalKickRight   PlayMode = "goal_kick_right"
	AfterGoalLeft   PlayMode = "after_goal_left"
	AfterGoalRight  PlayMode = "after_goal_right"
	DropBall        PlayMode = "drop_ball"
	OffSideLeft     PlayMode = "off_side_left"
	OffSideRight    PlayMode = "off_side_right"

	PkLeft                    PlayMode = "pk_left"
	PkRight                   PlayMode = "pk_right"
	FirstHalfOver             PlayMode = "first_half_over"
	Pause                     PlayMode = "pause"
	Human                     PlayMode = "human"
	FoulChargeLeft            PlayMode = "foul_charge_left"
	FoulChargeRight           PlayMode = "foul_charge_right"
	FoulPushLeft              PlayMode = "foul_push_left"
	FoulPushRight             PlayMode = "foul_push_right"
	FoulMultipleAttackerLeft  PlayMode = "foul_multiple_attacker_left"
	FoulMultipleAttackerRight PlayMode = "foul_multiple_attacker_right"
	FoulBallOutLeft           PlayMode = "foul_ball_out_left"
	FoulBallOutRight          PlayMode = "foul_ball_out_right"
	Max                       PlayMode = "max"
)

type Init struct {
	Side          Side
	UniformNumber UniformNumber
	PlayMode      PlayMode
}

func (m *Init) UnmarshalRcss(msg Message) error {
	if _, err := fmt.Sscanf(msg.values[0], "%c", &m.Side); err != nil {
		return fmt.Errorf("error on parsing side: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[1], &m.UniformNumber); err != nil {
		return fmt.Errorf("error on parsing unum %s", err)
	}

	if _, err := fmt.Sscan(msg.values[2], &m.PlayMode); err != nil {
		return fmt.Errorf("error on parsing mode %s", err)
	}

	return nil
}

type ServerParameters struct {
	// Goal width
	GoalWidth float32

	// Player size
	PlayerSize float32

	// Player decay
	PlayerDecay float32

	PlayerRand float32

	// Player weight
	PlayerWeight float32

	// Maximum player velocity
	PlayerSpeedMax float32

	// Maximum player acceleration
	PlayerAccelMax float32

	// Maximum player stamina
	StaminaMax float32

	// Maximum player stamina increment
	StaminaIncMax float32

	// Player recovery decrement threshold
	RecoverDecThr float32

	// Minimum player recovery
	RecoverMin float32

	// Player recovery decrement
	RecoverDec float32

	// player dash effort decrement threshold
	EffortDecThr float32
}

func (m *ServerParameters) UnmarshalRcss(msg Message) error {
	return fmt.Errorf("not implemented")
}

type PlayerParameters struct {
}

func (m *PlayerParameters) UnmarshalRcss(msg Message) error {
	return fmt.Errorf("not implemented")
}

type PlayerType struct {
	ID int
}

func (m *PlayerType) UnmarshalRcss(msg Message) error {
	return fmt.Errorf("not implemented")
}

// Input Driver
type Team interface {
	Name() string

	Init(match Match, side Side, unum UniformNumber, mode PlayMode)
	ServerParam(sp ServerParameters)
	PlayerParam(pp PlayerParameters)
	PlayerType(pt PlayerType)

	See()
	Hear()
	SenseBody()

	Score()
}
