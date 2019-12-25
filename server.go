package rcss

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	chewxySexp "github.com/chewxy/chexySexp"
	"github.com/nsf/sexp"
)

type Server interface {
	Match

	Stop() error
}

type server struct {
	raddr *net.UDPAddr
	conn  net.PacketConn
}

// type Init struct {
// 	init []string `sexp:"init,siblings"`
// }

func NewServer(addr string) (Server, error) {
	if raddr, err := net.ResolveUDPAddr("udp", addr); err != nil {
		return nil, err
	} else if conn, err := createUdpConn(); err != nil {
		return nil, err
	} else {
		srv := &server{
			raddr: raddr,
			conn:  conn,
		}

		return srv, nil
	}
}

func (s server) Stop() error {
	return s.conn.Close()
}

func (s server) bind(team Team) {
	l := make([]byte, 4096)
	var str string
	var ast *sexp.Node
	for {
		if n, _, err := s.conn.ReadFrom(l); err != nil {
			fmt.Printf("error: %s\n", err)
			return
		} else {
			//saw := false
			str = string(l[:n])
			ast, err = sexp.Parse(strings.NewReader(str), nil)
			if nil != err {
				fmt.Printf("error on : %s\n", err)
			}
			// var msg Message
			// if err := msg.UnmarshalBinary(l); err != nil {
			// 	fmt.Printf("message parse error: %s\n", err)
			// 	continue
			// }

			//fmt.Printf("%#v\n", msg)
			Sexp, err := chewxySexp.ParseString(str)
			if nil != err {
				fmt.Printf("Error : %s\n", err)
			}

			switch fmt.Sprint(Sexp[0].Head()) {
			case "init":
				var m Init
				err := ast.Unmarshal(&m.Init)
				if nil != err {
					panic(err)
				} else {
					m.SetValues()
					go ProcessInit(s, m, team)
				}

				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal Init message: %s\n", err)

				// 	continue
				// }

				//go team.Init(s, m.Side, m.UniformNumber, m.PlayMode)
			case "server_param":

				var m ServerParameters

				child := fmt.Sprint(fmt.Sprint(Sexp[0].Tail()))
				child = child[1 : len(child)-1]

				newast, err := sexp.Parse(strings.NewReader(child), nil)
				if nil != err {
					fmt.Printf("Error on parsing server_param : %s\n", err)
				} else {
					newast.Unmarshal(&m)
					go ProcessServerParam(m, team)
					// if nil != err {
					// 	fmt.Printf("Error on unmarshaling server_param : %s\n", err)
					// }
				}

				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal ServerParameters message: %s\n", err)

				// 	continue
				// }

				// go team.ServerParam(m)
				// err := ast.Unmarshal(&(m.ServerParameters))
				// if nil != err {
				// 	fmt.Printf("error : %s\n", err)
				// } else {
				// 	fmt.Println("vals")
				// 	m.SetValues()
				// }

			case "player_param":
				var m PlayerParameters

				child := fmt.Sprint(fmt.Sprint(Sexp[0].Tail()))
				child = child[1 : len(child)-1]

				newast, err := sexp.Parse(strings.NewReader(child), nil)
				if nil != err {
					fmt.Printf("Error on parsing player_param : %s\n", err)
				} else {
					newast.Unmarshal(&m)
					go ProcessPlayerParam(m, team)
					// if nil != err {
					// 	fmt.Printf("Error on unmarshaling player_param : %s\n", err)

					// }
				}

				//fmt.Println(fmt.Sprint(Sexp[0].Head()))

				// if err := m.UnmarshalRcss(msg); err != nil {
				// 	log.Printf("error on unmarshal PlayerParameters message: %s\n", err)

				// 	continue
				// }

				// go team.PlayerParam(m)

			case "player_type":
				var m PlayerType

				child := fmt.Sprint(fmt.Sprint(Sexp[0].Tail()))
				child = child[1 : len(child)-1]

				newast, err := sexp.Parse(strings.NewReader(child), nil)
				if nil != err {
					fmt.Printf("Error on parsing player_type : %s\n", err)
				} else {
					newast.Unmarshal(&m)
					go ProcessPlayerType(m, team)
					// if nil != err {
					// 	fmt.Printf("Error on unmarshaling player_type : %s\n", err)

					// }
				}

			case "see":
				var m See
				SeperateSeeParam(&m, str)
				for i, v := range m.see.Array {
					if i == 0 {
						continue
					}
					obj := ProcessSee(v)
					if obj.Head() == "b" {
						fmt.Println(v)
						fmt.Println(obj)
					}

				}
				// var m See
				// child := fmt.Sprint(fmt.Sprint(Sexp[0].Tail()))
				// child = child[1 : len(child)-1]

				// newast, err := sexp.Parse(strings.NewReader(child), nil)
				// if nil != err {
				// 	fmt.Println("Error", err)
				// } else {
				// 	newast.Unmarshal(&m.see)
				// 	if false == saw {
				// 		fmt.Println("Array : ", m.see.Array)
				// 		saw = true
				// 	}
				// }

			case "hear":
				fmt.Println("Hear")
				var m Hear
				err := ast.Unmarshal(&m.Hear)
				if nil != err {
					panic(err)
				} else {
					m.SetValues()
					fmt.Println(m)
					ProcessHear(m, team)

				}

			case "sense_body":

			case "score":

			case "error":

			default:
				//fmt.Printf("unhandled server input: `%s`\n", str)
			}
		}
	}
	//Should be ommited , I just put it to get rid og "ast declered and not used" error!
	//ast.Unmarshal()
}

func newInitCommand(teamName string, goalie bool, version int) Message {
	msg := Message{name: "init"}

	msg.AddValues(teamName)
	if version > 0 {
		ver := Message{name: "version"}
		ver.AddValues(strconv.Itoa(version))
		msg.AddSubmessages(ver)
	}
	if goalie {
		g := Message{name: "goalie"}
		msg.AddSubmessages(g)
	}

	return msg
}

func (s server) Join(team Team) error {
	go s.bind(team)

	cmd := newInitCommand(team.Name(), false, 7)

	if b, err := cmd.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Reconnect(team Team, unum UniformNumber) error {
	msg := Message{name: "reconnect"}
	msg.AddValues(team.Name(), strconv.Itoa(int(unum)))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}

}

func (s server) Bye() error {
	msg := Message{name: "bye"}
	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
	return nil
}

func (s server) Catch(dir Direction) error {
	msg := Message{name: "catch"}
	msg.AddValues(strconv.FormatFloat(float64(dir), 'f', -1, 64))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) ChangeView(w SightWidth, q SightQuality) error {
	msg := Message{name: "change_view"}
	msg.AddValues(string(w), string(q))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Dash(p Power) error {
	msg := Message{name: "dash"}
	msg.AddValues(strconv.Itoa(int(p)))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Kick(p Power, d Direction) error {
	msg := Message{name: "kick"}
	msg.AddValues(strconv.Itoa(int(p)), strconv.FormatFloat(float64(d), 'f', -1, 64))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Move(x, y int) error {
	n, err := s.conn.WriteTo([]byte(fmt.Sprintf("(move %d %d)", x, y)), s.raddr)

	fmt.Println(n, err)

	return err
}

func (s server) Say(msg string) error {
	_, err := s.conn.WriteTo([]byte(msg), s.raddr)

	return err
}

func (s server) Turn(m Moment) error {
	msg := Message{name: "turn"}
	msg.AddValues(strconv.Itoa(int(m)))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) TurnNeck(n NeckMoment) error {
	msg := Message{name: "turn_neck"}
	msg.AddValues(strconv.Itoa(int(n)))

	if b, err := msg.MarshalBinary(); err != nil {
		return err
	} else if n, err := s.conn.WriteTo(b, s.raddr); err != nil {
		return err
	} else if 0 == n {
		return fmt.Errorf("nothing has been written")
	} else {
		return nil
	}
}

func (s server) Score() error {
	_, err := s.conn.WriteTo([]byte("(score)"), s.raddr)
	return err
}

func (s server) See() error {
	_, err := s.conn.WriteTo([]byte("(see)"), s.raddr)
	return err
}

func (s server) SenseBody() error {
	_, err := s.conn.WriteTo([]byte("(sense_body)"), s.raddr)
	return err
}
