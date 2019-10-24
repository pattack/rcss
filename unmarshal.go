package rcss

func ProcessInit(match Match, msg Init, team Team) {
	team.SetSide(msg.Side)
	team.SetPlayMode(msg.PlayMode)
	team.Invite(match, msg.UniformNumber)
}
func ProcessServerParam(m ServerParameters, team Team) {
	team.ServerParam(m)
}
func ProcessPlayerParam(m PlayerParameters, team Team) {
	team.PlayerParam(m)
}
func ProcessPlayerType(m PlayerType, team Team) {
	team.PlayerType(m)
}
func ProcessHear(m Hear, team Team) {
	if PlayMode(m.Message) == KickOffLeft || PlayMode(m.Message) == KickOffRight {
		team.Kickoff()
	}

}
