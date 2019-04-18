package rcss

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type Server interface {
	Match

	Stop() error
}

type server struct {
	raddr *net.UDPAddr
	conn  net.PacketConn
}

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

	for {
		if _, _, err := s.conn.ReadFrom(l); err != nil {
			fmt.Printf("error: %s\n", err)

			return
		} else {
			var msg Message
			if err := msg.UnmarshalBinary(l); err != nil {
				fmt.Printf("message parse error: %s\n", err)

				continue
			}

			//fmt.Printf("%#v\n", msg)

			switch msg.name {
			case "init":
				var m Init
				if err := m.UnmarshalRcss(msg); err != nil {
					log.Printf("error on unmarshal Init message: %s\n", err)

					continue
				}

				go team.Init(s, m.Side, m.UniformNumber, m.PlayMode)

			case "server_param":
				var m ServerParameters
				if err := m.UnmarshalRcss(msg); err != nil {
					log.Printf("error on unmarshal ServerParameters message: %s\n", err)

					continue
				}

				go team.ServerParam(m)

			case "player_param":
				var m PlayerParameters
				if err := m.UnmarshalRcss(msg); err != nil {
					log.Printf("error on unmarshal PlayerParameters message: %s\n", err)

					continue
				}

				go team.PlayerParam(m)

			case "player_type":
				var m PlayerType
				if err := m.UnmarshalRcss(msg); err != nil {
					log.Printf("error on unmarshal PlayerType message: %s\n", err)

					continue
				}

				go team.PlayerType(m)

			case "see":

			case "hear":

			case "sense_body":

			case "score":

			case "error":

			default:
				fmt.Printf("unhandled server input: `%s`\n", msg.name)
			}
		}
	}
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

	cmd := newInitCommand(team.Name(), false, 15)

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
	return nil
}

func (s server) Bye() error {
	return nil
}

func (s server) Catch(dir Direction) error {
	return nil
}

func (s server) ChangeView(w SightWidth, q SightQuality) error {
	return nil
}

func (s server) Dash() error {
	return nil
}

func (s server) Kick() error {
	return nil
}

func (s server) Move(x, y int) error {
	_, err := s.conn.WriteTo([]byte(fmt.Sprintf("(move %d %d)", x, y)), s.raddr)

	return err
}

func (s server) Say() error {
	return nil
}

func (s server) Turn() error {
	return nil
}

func (s server) TurnNeck() error {
	return nil
}

func (s server) Score() error {
	return nil
}

func (s server) See() error {
	return nil
}

func (s server) SenseBody() error {
	return nil
}
