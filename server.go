package rcss

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

// Output Driver
type Match interface {
	Join(team Team) error
	Reconnect(team Team, unum UniformNumber) error
	Bye() error

	Catch() error
	ChangeView() error
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

type Server interface {
	Match

	Stop() error
}

type server struct {
	raddr *net.UDPAddr
	conn net.PacketConn
}

func NewServer(addr string) (Server, error) {
	if raddr, err := net.ResolveUDPAddr("udp", addr); err != nil {
		return nil, err
	} else if conn, err := createUdpConn(); err != nil {
		return nil, err
	} else {
		srv := &server{
			raddr: raddr,
			conn: conn,
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
			scn := bufio.NewScanner(bytes.NewReader(l))
			scn.Split(bufio.ScanWords)

			if scn.Scan() {
				switch scn.Text() {
				case "(init":
					scn.Scan()
					scn.Text()
					side := LeftSide

					scn.Scan()
					scn.Text()
					unum := UniformNumber(1)

					scn.Scan()
					scn.Text()
					mode := BeforeKickOff

					team.Init(s, side, unum, mode)

				case "(see":

				case "(sense_body":

				case "(error":

				default:
					fmt.Printf("unhandled server input: `%s`\n", scn.Text())
				}
			}
		}
	}
}

func (s server) Join(team Team) error {
	go s.bind(team)

	_, err := s.conn.WriteTo([]byte(fmt.Sprintf("(init %s (version 15))", team.Name())), s.raddr)

	return err
}

func (s server) Reconnect(team Team, unum UniformNumber) error {
	return nil
}

func (s server) Bye() error {
	return nil
}

func (s server) Catch() error {
	return nil
}

func (s server) ChangeView() error {
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
