package rcss

import (
	"fmt"

	chewxySexp "github.com/chewxy/chexySexp"
)

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
func SeperateSeeParam(m *See, str string) {
	Sexp, err := chewxySexp.ParseString(str)
	if err != nil {
		fmt.Println("See Error : ", err)
	}
	child := fmt.Sprint(Sexp[0].Tail())
	Sexp, err = chewxySexp.ParseString(child)
	for child != "()" {
		head := fmt.Sprint(Sexp[0].Head())
		fmt.Println(head)
		m.see.Array = append(m.see.Array, head)
		child = fmt.Sprint(Sexp[0].Tail())
		fmt.Println(child)
		Sexp, err = chewxySexp.ParseString(child)

	}
	fmt.Println(m.see.Array)
}
func ProcessSee(obj string) {
	Sexp, err := chewxySexp.ParseString(obj)
	if nil != err {
		fmt.Printf("Error : %s\n", err)
	}
	head := fmt.Sprint(fmt.Sprint(Sexp[0].Head()))
	Sexp, err = chewxySexp.ParseString(head)
	if nil != err {
		panic(err)
	}
	head = fmt.Sprint(Sexp[0].Head())

	if head == "f" {
		fmt.Println("flag")
	} else {
		fmt.Println(head)
	}
}
