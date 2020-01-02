package rcss

import (
	"fmt"
	"strconv"

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
		m.see.Array = append(m.see.Array, head)
		child = fmt.Sprint(Sexp[0].Tail())
		Sexp, err = chewxySexp.ParseString(child)

	}

}
func ProcessFlags(obj string, time int) Flag {
	var f Flag
	f.Set()
	f.Time = time
	Sexp, err := ToSexp(obj)
	if err != nil {
		fmt.Println("Obj Error : ", err)
	}

	child := SexpTail(Sexp)
	Head := SexpHead(Sexp)
	Sexp, err = ToSexp(Head)
	Head = SexpTail(Sexp)
	Sexp, err = ToSexp(Head)
	for Head != "()" && Head != "<nil>" {
		head := SexpHead(Sexp)
		if head == "l" {
			f.Left = true
		} else if head == "r" {
			f.Right = true
		} else if head == "t" {
			f.Top = true
		} else if head == "b" {
			f.Bottom = true
		} else if head == "c" {
			f.Center = true
		} else if head == "g" {
			f.Goal = true
		} else if head == "p" {
			f.Penalty = true
		} else {
			f.Number, err = strconv.ParseFloat(head, 64)
			if nil != err {
				fmt.Println("Error on parsing string to int : ", err)
			}
		}
		Head = SexpTail(Sexp)
		Sexp, err = ToSexp(Head)

	}
	Sexp, err = chewxySexp.ParseString(child)
	Dis := SexpHead(Sexp)
	Dir := fmt.Sprint(Sexp[0].Tail().Head())
	f.Dis, _ = strconv.ParseFloat(Dis, 64)
	f.Dir, _ = strconv.ParseFloat(Dir, 64)
	return f
}
func ProcessGoals(obj string, time int) Goal {
	var g Goal
	g.Set()
	g.Time = time
	Sexp, err := ToSexp(obj)
	if err != nil {
		fmt.Println("Obj Error : ", err)
	}

	child := SexpTail(Sexp)
	Head := SexpHead(Sexp)
	Sexp, err = ToSexp(Head)
	Head = fmt.Sprint(Sexp[0].Tail().Head())
	if Head == "l" {
		g.Left = true
	} else {
		g.Right = true
	}
	Sexp, err = chewxySexp.ParseString(child)
	Dis := SexpHead(Sexp)
	Dir := fmt.Sprint(Sexp[0].Tail().Head())
	g.Dis, _ = strconv.ParseFloat(Dis, 64)
	g.Dir, _ = strconv.ParseFloat(Dir, 64)
	return g

}
func ProcessBall(obj string, time int) Ball {
	var b Ball
	b.Set()
	b.Time = time
	Sexp, err := ToSexp(obj)
	if err != nil {
		fmt.Println("Obj Error : ", err)
	}
	var Datas []float64
	child := SexpTail(Sexp)

	for child != "()" && child != "<nil>" {
		head := SexpHeadString(child)
		x, _ := strconv.ParseFloat(head, 64)
		Datas = append(Datas, x)
		child = SexpTailString(child)
	}
	b.Dis = Datas[0]
	b.Dir = Datas[1]
	if len(Datas) == 4 {

		b.DisChng = Datas[2]
		b.DirChng = Datas[3]
	}

	return b

}

func ProcessLine(obj string, time int) Line {
	var l Line
	l.Set()
	l.Time = time
	Sexp, err := ToSexp(obj)
	if err != nil {
		fmt.Println("Obj Error : ", err)
	}

	child := SexpTail(Sexp)
	Head := SexpHead(Sexp)
	Head = SexpTailString(Head)
	Head = SexpHeadString(Head)
	if Head == "l" {
		l.Left = true
	} else if Head == "r" {
		l.Right = true
	} else if Head == "t" {
		l.Top = true
	} else if Head == "b" {
		l.Bottom = true
	}
	Dis := SexpHeadString(child)
	Dir := SexpHeadString(SexpTailString(child))
	l.Dis, _ = strconv.ParseFloat(Dis, 64)
	l.Dir, _ = strconv.ParseFloat(Dir, 64)
	return l
}

func ProcessSee(obj string, time string) Object {
	Time, _ := strconv.Atoi(time)
	Sexp, err := ToSexp(obj)
	if err != nil {
		fmt.Println("Obj Error : ", err)
	}
	Head := SexpHead(Sexp)
	if Head == "b" {
		return ProcessBall(obj, Time)
	} else if Head == "F" {
		fmt.Println("UnSuported F")
	} else {
		Sexp, err = ToSexp(Head)
		if Type := SexpHead(Sexp); Type == "f" {
			return ProcessFlags(obj, Time)
		} else if Type == "g" {
			return ProcessGoals(obj, Time)
		} else if Type == "l" {
			return ProcessLine(obj, Time)
		}
	}

	return Flag{}
}
func ToSexp(text string) ([]chewxySexp.Sexp, error) {
	Sexp, err := chewxySexp.ParseString(text)
	return Sexp, err
}
func SexpHead(Sexp []chewxySexp.Sexp) string {
	return fmt.Sprint(Sexp[0].Head())
}
func SexpTail(Sexp []chewxySexp.Sexp) string {
	return fmt.Sprint(Sexp[0].Tail())
}
func SexpHeadString(text string) string {
	Sexp, _ := ToSexp(text)
	return SexpHead(Sexp)
}
func SexpTailString(text string) string {
	Sexp, _ := ToSexp(text)
	return SexpTail(Sexp)
}
