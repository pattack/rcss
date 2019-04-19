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

func (m Init) adapter() *buffer {
	return &buffer{
		name: "init",
		vars: {
			{title: "side", format: "%c", value: m.Side},
			{title: "unum", value: m.UniformNumber},
			{title: "mode", value: m.PlayMode},
		},
	}
}

func (m *Init) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *Init) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

// Aggregate of 109 Server Parameters
type ServerParameters struct {
	// Goal width
	// name: goal_width
	GoalWidth float32

	// Player size
	// name: player_size
	PlayerSize float32

	// Player decay
	// name: player_decay
	PlayerDecay float32

	//
	// name: player_rand
	PlayerRand float32

	// Player weight
	// name: player_weight
	PlayerWeight float32

	// Maximum player velocity
	// name: player_speed_max
	MaxPlayerSpeed float32

	// Maximum player acceleration
	// name: player_accel_max
	MaxPlayerAcceleration float32

	// Maximum player stamina
	// name: stamina_max
	MaxStamina float32

	// Maximum player stamina increment
	// name: stamina_inc_max
	MaxStaminaIncrement float32

	// Player recovery decrement threshold
	// name: recover_dec_thr
	PlayerRecoveryDecrementThreshold float32

	// Minimum player recovery
	// name: recover_min
	MinPlayerRecovery float32

	// Player recovery decrement
	// name: recover_dec
	PlayerRecoveryDecrement float32

	// Player dash effort decrement threshold
	// name: effort_dec_thr
	EffortDecrementThreshold float32

	// Minimum player dash effort
	// name: effort_min
	MinEffort float32

	// Player dash effort decrement
	// name: effort_dec
	EffortDecrement float32

	// Dash effort increment threshold
	// name: effort_inc_thr
	EffortIncrementThreshold float32

	// Dash effort increment
	// name: effort_inc
	EffortIncrement float32

	// Noise added directly to kicks
	// name: kick_rand
	KickRand float32

	// Flag whether to use team specific actuator noise
	// name: team_actuator_noise
	TeamActuatorNoise bool

	// Factor to multiply prand for left team
	// name: prand_factor_l
	LeftPlayerRandFactor float32

	// Factor to multiply prand for right team
	// name: prand_factor_r
	RightPlayerRandFactor float32

	// Factor to multiply kick rand for left team
	// name: kick_rand_factor_l
	LeftKickRandFactor float32

	// Factor to multiply kick rand for right team
	// name: kick_rand_factor_r
	RightKickRandFactor float32

	// Ball size
	// name: ball_size
	BallSize float32

	// Ball decay
	// name: ball_decay
	BallDecay float32

	//
	// name: ball_rand
	BallRand float32

	// Weight of the ball
	// name: ball_weight
	BallWeight float32

	// Maximum ball velocity
	// name: ball_speed_max
	MaxBallSpeed float32

	// Maximum ball acceleration
	// name: ball_accel_max
	MaxBallAcceleration float32

	// Dash power rate
	// name: dash_power_rate
	DashPowerRate float32

	//
	// name: kick_power_rate
	KickPowerRate float32

	// Kickable margin
	// name: kickable_margin
	KickableMargin float32

	// Control radius
	// name: control_radius
	ControlRadius float32

	// Goalie catch probability
	// name: catch_probability
	GoalieCatchProbability float32

	// Goalie catchable area length
	// name: catchable_area_l
	GoalieCatchableAreaLength float32

	// Goalie catchable area width
	// name: catchable_area_w
	GoalieCatchableAreaWidth float32

	// Goalie maximum moves after a catch
	// name: goalie_max_moves
	MaxGoalieAfterCatchMoves int

	// Maximum power
	// name: maxpower
	MaxPower int

	// Minumum power
	// name: minpower
	MinPower int

	// Maximum moment
	// name: maxmoment
	MaxMoment int

	// Minimum moment
	// name: minmoment
	MinMoment int

	// Maximum neck moment
	// name: maxneckmoment
	MaxNeckMoment int

	// Minimum neck moment
	// name: minneckmoment
	MinNeckMoment int

	// Maximum neck angle
	// name: maxneckang
	MaxNeckAngle int

	// Minimum neck angle
	// name: minneckang
	MinNeckAngle int

	// Visible angle
	// name: visible_angle
	VisibleAngle float32

	// Visible distance
	// name: visible_distance
	VisibleDistance float32

	// Audio cut off distance
	// name: audio_cut_dist
	AudioCutOffDistance float32

	// Quantize step of distance for movable objects
	// name: quantize_step
	MovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for landmarks
	// name: quantize_step_l
	LandmarksDistanceQuantizeStep float32

	// Quantize step of direction
	// name: quantize_step_dir
	DirectionQuantizeStep float32

	// Quantize step of distance for movable objects for left team
	// name: quantize_step_dist_team_l
	LeftTeamMovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for movable objects for right team
	// name: quantize_step_dist_team_r
	RightTeamMovableObjectsDistanceQuantizeStep float32

	// Quantize step of distance for landmarks for left team
	// name: quantize_step_dist_l_team_l
	LeftTeamLandmarksDistanceQuantizeStep float32

	// Quantize step of distance for landmarks for right team
	// name: quantize_step_dist_l_team_r
	RightTeamLandmarksDistanceQuantizeStep float32

	// Quantize step of direction for left team
	// name: quantize_step_dir_team_l
	LeftTeamDirectionQuantizeStep float32

	// Quantize step of direction for right team
	// name: quantize_step_dir_team_r
	RightTeamDirectionQuantizeStep float32

	// Corner Kick Margin
	// name: ckick_margin
	CornerKickMargin float32

	// Wind direction
	// name: wind_dir
	WindDirection float32

	//
	// name: wind_force
	WindForce float32

	//
	// name: wind_rand
	WindRand float32

	// Wind factor is none
	// name: wind_none
	NoWind bool

	// Wind factor is random
	// name: wind_random
	ProbableWind bool

	// Inertia moment for turn
	// name: inertia_moment
	InertiaMoment float32

	// Length of a half time in seconds
	// name: half_time
	// TODO: use time.Duration
	HalfTime int

	// Number of cycles to wait until dropping the ball automatically
	// name: drop_ball_time
	DropBallTime int

	// Player port number
	// name: port
	Port int

	// Offline coach port
	// name: coach_port
	OfflineCoachPort int

	// Online coach port
	// name: olcoach_port
	OnlineCoachPort int

	// Upper limit of the number of online coach’s message
	// name: say_coach_cnt_max
	OnlineCoachMaxMessageNumber int

	// Upper limit of length of online coach’s message
	// name: say_coach_msg_size
	OnlineCoachMaxMessageLength int

	// Time step of simulation [unit:msec]
	// name: simulator_step
	SimulatorStep int

	// Time step of visual information [unit:msec]
	// name: send_step
	SendStep int

	// Time step of acception of commands [unit: msec]
	// name: recv_step
	ReceiveStep int

	// Time step of body being sensed
	// name: sense_body_step
	SenseBodyStep int

	// String size of say message [unit:byte]
	// name: say_msg_size
	SayMessageSize int

	// Time window which controls how many messages can be sent (coach language)
	// name: clang_win_size
	CoachLanguageWindowSize int

	// Number of messages per window
	// name: clang_define_win
	CoachLanguageMessagesPerWindow int

	//
	// name: clang_meta_win
	CoachLanguageMetaWindow int

	//
	// name: clang_advice_win
	CoachLanguageAdviceWindow int

	//
	// name: clang_info_win
	CoachLanguageInformationWindow int

	// Delay between receipt of message and send to players
	// name: clang_mess_delay
	CoachLanguageMessageDelay int

	// Maximum number of coach messages sent per cycle
	// name: clang_mess_per_cycle
	CoachLanguageMaxMessagesPerCycle int

	//
	// name: head_max
	MaxHear int

	//
	// name: hear_inc
	HearIncrement int

	//
	// name: hear_decay
	HearDecay int

	//
	// name: catch_ban_cycle
	CatchBanCycle int

	// ?
	// name: coach
	Coach bool

	// ?
	// name: coach_w_referee
	CoachWithReferee bool

	// ?
	// name: old_coach_hear
	OldCoachHear bool

	// Interval of online coach’s look
	// name: send_vi_step
	OnlineCoachLookInterval int

	// Flag for using off side rule
	// name: use_offside
	UseOffside bool

	// Offside active area size
	// name: offside_active_area_size
	OffsideActiveAreaSize int

	// forbid kick off offside
	// name: forbid_kick_off_offside
	ForbidKickOffOffside bool

	//
	// name: log_file
	LogFile string

	//
	// name: record
	Record bool

	//
	// name: record_version
	RecordVersion int

	// Flag for record log
	// name: record_log
	RecordLog bool

	// Flag for record client command log
	// name: record_messages
	RecordMessages bool

	// Flag for send client command log
	// name: send_log
	SendLog bool

	// Flag for writing cycle lenth to log file
	// name: log_times
	LogTimes bool

	// Flag for verbose mode
	// name: verbose
	Verbose bool

	//
	// name: replay
	Replay bool

	// Offsie kick margin
	// name: offside_kick_margin
	OffsideKickMargin float32

	//
	// name: slow_down_factor
	SlowDownFactor float32

	// ?
	// name: start_goal_l
	LeftStartGoal string

	// ?
	// name: start_goal_r
	RightStartGoal string

	// ?
	// name: fullstate_l
	LeftFullState string

	// ?
	// name: fullstate_r
	RightFullState string
}

func (m ServerParameters) adapter() *buffer {
	var temp string

	return &buffer{
		name: "server_param",

		vars: {
			{name: "gwidth", m.GoalWidth},
			{name: "inertia_moment", m.InertiaMoment},
			{name: "psize", m.PlayerSize},
			{name: "pdecay", m.PlayerDecay},
			{name: "prand", m.PlayerRand},
			{name: "pweight", m.PlayerWeight},
			{name: "pspeed_max", m.MaxPlayerSpeed},
			{name: "paccel_max", m.MaxPlayerAcceleration},
			{name: "stamina_max", m.MaxStamina},
			{name: "stamina_inc", m.MaxStaminaIncrement},

			{name: "recover_init", temp},
			{name: "recover_dthr", m.PlayerRecoveryDecrementThreshold},
			{name: "recover_min", temp},
			{name: "recover_dec", m.PlayerRecoveryDecrement},
			{name: "effort_init", temp},
			{name: "effort_dthr", m.EffortDecrementThreshold},
			{name: "effort_min", m.MinEffort},
			{name: "effort_dec", m.EffortDecrement},
			{name: "effort_ithr", m.EffortIncrementThreshold},
			{name: "effort_inc", m.EffortIncrement},

			{name: "kick_rand", m.KickRand},
			{name: "team_actuator_noise", m.TeamActuatorNoise},
			{name: "prand_factor_l", m.LeftPlayerRandFactor},
			{name: "prand_factor_r", m.RightPlayerRandFactor},
			{name: "kick_rand_factor_l", m.LeftKickRandFactor},
			{name: "kick_rand_factor_r", m.RightKickRandFactor},
			{name: "bsize", m.BallSize},
			{name: "bdecay", m.BallDecay},
			{name: "brand", m.BallRand},
			{name: "bweight", m.BallWeight},

			{name: "bspeed_max", m.MaxBallSpeed},
			{name: "baccel_max", m.MaxBallAcceleration},
			{name: "dprate", m.DashPowerRate},
			{name: "kprate", m.KickPowerRate},
			{name: "kmargin", m.KickableMargin},
			{name: "ctlradius", m.ControlRadius},
			{name: "ctlradius_width", temp},
			{name: "maxp", m.MaxPower},
			{name: "minp", m.MinPower},
			{name: "maxm", m.MaxMoment},

			{name: "minm", m.MinMoment},
			{name: "maxnm", m.MaxNeckMoment},
			{name: "minnm", m.MinNeckMoment},
			{name: "maxn", m.MaxNeckAngle},
			{name: "minn", m.MaxNeckAngle},
			{name: "visangle", m.VisibleAngle},
			{name: "visdist", m.VisibleDistance},
			{name: "windir", m.WindDirection},
			{name: "winforce", m.WindForce},
			{name: "winang", temp},

			{name: "winrand", m.WindRand},
			{name: "kickable_area", temp},
			{name: "catch_area_l", m.GoalieCatchableAreaLength},
			{name: "catch_area_w", m.GoalieCatchableAreaWidth},
			{name: "catch_prob", m.GoalieCatchProbability},
			{name: "goalie_max_moves", m.MaxGoalieAfterCatchMoves},
			{name: "ckmargin", m.CornerKickMargin},
			{name: "offside_area", m.OffsideActiveAreaSize},
			{name: "win_no", m.NoWind},
			{name: "win_random", m.ProbableWind},

			{name: "say_cnt_max", temp},
			{name: "SayCoachMsgSize", temp},
			{name: "clang_win_size", m.CoachLanguageWindowSize},
			{name: "clang_define_win", temp},
			{name: "clang_meta_win", m.CoachLanguageMetaWindow},
			{name: "clang_advice_win", m.CoachLanguageAdviceWindow},
			{name: "clang_info_win", m.CoachLanguageInformationWindow},
			{name: "clang_mess_delay", m.CoachLanguageMessageDelay},
			{name: "clang_mess_per_cycle", m.CoachLanguageMaxMessagesPerCycle},
			{name: "half_time", m.HalfTime},

			{name: "sim_st", m.SimulatorStep},
			{name: "send_st", m.SendStep},
			{name: "recv_st", m.ReceiveStep},
			{name: "sb_step", m.SenseBodyStep},
			{name: "lcm_st", temp},
			{name: "SayMsgSize", m.SayMessageSize},
			{name: "hear_max", m.MaxHear},
			{name: "hear_inc", m.HearIncrement},
			{name: "hear_decay", m.HearDecay},
			{name: "cban_cycle", m.CatchBanCycle},

			{name: "slow_down_factor", m.SlowDownFactor},
			{name: "useoffside", m.UseOffside},
			{name: "kickoffoffside", temp},
			{name: "offside_kick_margin", m.OffsideKickMargin},
			{name: "audio_dist", m.AudioCutOffDistance},
			{name: "dist_qstep", m.MovableObjectsDistanceQuantizeStep},
			{name: "land_qstep", m.LandmarksDistanceQuantizeStep},
			{name: "dir_qstep", m.DirectionQuantizeStep},
			{name: "dist_qstep_l", m.LeftTeamMovableObjectsDistanceQuantizeStep},
			{name: "dist_qstep_r", m.RightTeamMovableObjectsDistanceQuantizeStep},

			{name: "land_qstep_l", m.LeftTeamLandmarksDistanceQuantizeStep},
			{name: "land_qstep_r", m.RightTeamLandmarksDistanceQuantizeStep},
			{name: "dir_qstep_l", m.LeftTeamDirectionQuantizeStep},
			{name: "dir_qstep_r", m.RightTeamDirectionQuantizeStep},
			{name: "CoachMode", m.Coach},
			{name: "CwRMode", m.CoachWithReferee},
			{name: "old_hear", m.OldCoachHear},
			{name: "sv_start", temp},
			{name: "start_goal_l", m.LeftStartGoal},
			{name: "start_goal_r", m.RightStartGoal},

			{name: "fullstate_l", m.LeftFullState},
			{name: "fullstate_r", m.RightFullState},
			{name: "drop_time", m.DropBallTime},
			{name: "", temp},
			{name: "", temp},
			{name: "", temp},
			{name: "", temp},
			{name: "", temp},
			{name: "", temp},
		},
	}
}

func (m *ServerParameters) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *ServerParameters) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

type PlayerParameters struct {
	PlayerTypes int

	SubsMax int

	PtMax int

	PlayerSpeedMaxDeltaMin   int
	PlayerSpeedMaxDeltaMax   int
	StaminaIncMaxDeltaFactor int

	PlayerDecayDeltaMin      int
	PlayerDecayDeltaMax      int
	InertiaMomentDeltaFactor int

	DashPowerRateDeltaMin int
	DashPowerRateDeltaMax int
	PlayerSizeDeltaFactor int

	KickableMarginDeltaMin int
	KickableMarginDeltaMax int
	KickRandDeltaFactor    int

	ExtraStaminaDeltaMin int
	ExtraStaminaDeltaMax int
	EffortMaxDeltaFactor int
	EffortMinDeltaFactor int

	SpareLong1  int
	SpareLong2  int
	SpareLong3  int
	SpareLong4  int
	SpareLong5  int
	SpareLong6  int
	SpareLong7  int
	SpareLong8  int
	SpareLong9  int
	SpareLong10 int

	SpareShort1  int
	SpareShort2  int
	SpareShort3  int
	SpareShort4  int
	SpareShort5  int
	SpareShort6  int
	SpareShort7  int
	SpareShort8  int
	SpareShort9  int
	SpareShort10 int
}

func (m PlayerParameters) adapter() *buffer {
	return &rcssRecipe{
		{name: "player_types", value: m.PlayerTypes},
		{name: "subs_max", value: m.SubsMax},
		{name: "pt_max", value: m.PtMax},
		{name: "player_speed_max_delta_min", value: m.PlayerSpeedMaxDeltaMin},
		{name: "player_speed_max_delta_max", value: m.PlayerSpeedMaxDeltaMax},
		{name: "stamina_inc_max_delta_factor", value: m.StaminaIncMaxDeltaFactor},
		{name: "player_decay_delta_min", value: m.PlayerDecayDeltaMin},
		{name: "player_decay_delta_max", value: m.PlayerDecayDeltaMax},
		{name: "inertia_moment_delta_factor", value: m.InertiaMomentDeltaFactor},
		{name: "dash_power_rate_delta_min", value: m.DashPowerRateDeltaMin},
		{name: "dash_power_rate_delta_max", value: m.DashPowerRateDeltaMax},
		{name: "player_size_delta_factor", value: m.PlayerSizeDeltaFactor},
		{name: "kickable_margin_delta_min", value: m.KickableMarginDeltaMin},
		{name: "kickable_margin_delta_max", value: m.KickableMarginDeltaMax},
		{name: "kick_rand_delta_factor", value: m.KickRandDeltaFactor},
		{name: "extra_stamina_delta_min", value: m.ExtraStaminaDeltaMin},
		{name: "extra_stamina_delta_max", value: m.ExtraStaminaDeltaMax},
		{name: "effort_max_delta_factor", value: m.EffortMaxDeltaFactor},
		{name: "effort_min_delta_factor", value: m.EffortMinDeltaFactor},
	}
}

func (m *PlayerParameters) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *PlayerParameters) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
}

type PlayerType struct {
	// Player Identifier
	Id int

	// Maximum Player Speed
	PlayerSpeedMax int

	// Maximum Stamina Increase
	StaminaIncMax int

	// Player Decay
	PlayerDecay int

	// Inertia Moment
	InertiaMoment int

	// Dash Power Rate
	DashPowerRate int

	// Player Size
	PlayerSize int

	// Kickable Margin
	KickableMargin int

	// Kick Rand
	KickRand int

	// Extra Stamina
	ExtraStamina int

	// Maximum Effort
	EffortMax int

	// Minimum Effort
	EffortMin int

	SpareLong1  int
	SpareLong2  int
	SpareLong3  int
	SpareLong4  int
	SpareLong5  int
	SpareLong6  int
	SpareLong7  int
	SpareLong8  int
	SpareLong9  int
	SpareLong10 int
}

func (m *PlayerType) adapter() *buffer {
	return &buffer{}

	if _, err := fmt.Sscan(msg.values[0], &m.Id); err != nil {
		return fmt.Errorf("error on parsing id: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[1], &m.PlayerSpeedMax); err != nil {
		return fmt.Errorf("error on parsing player_speed_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[2], &m.StaminaIncMax); err != nil {
		return fmt.Errorf("error on parsing stamina_inc_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[3], &m.PlayerDecay); err != nil {
		return fmt.Errorf("error on parsing player_decay: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[4], &m.InertiaMoment); err != nil {
		return fmt.Errorf("error on parsing inertia_moment: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[5], &m.DashPowerRate); err != nil {
		return fmt.Errorf("error on parsing dash_power_rate: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[6], &m.PlayerSize); err != nil {
		return fmt.Errorf("error on parsing player_size: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[7], &m.KickableMargin); err != nil {
		return fmt.Errorf("error on parsing kickable_margin: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[8], &m.KickRand); err != nil {
		return fmt.Errorf("error on parsing kick_rand: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[9], &m.ExtraStamina); err != nil {
		return fmt.Errorf("error on parsing extra_stamina: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[10], &m.EffortMax); err != nil {
		return fmt.Errorf("error on parsing effort_max: %s", err)
	}

	if _, err := fmt.Sscan(msg.values[11], &m.EffortMin); err != nil {
		return fmt.Errorf("error on parsing effort_min: %s", err)
	}
}

func (m *PlayerType) UnmarshalRcss(msg Message) error {
	return m.adapter().UnmarshalRcss(msg)
}

func (m *PlayerType) MarshalRcss() (Message, error) {
	return m.adapter().MarshalRcss()
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
