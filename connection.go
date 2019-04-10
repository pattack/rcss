package rcss

import (
	"net"
)

func createUdpConn() (net.PacketConn, error) {
	if conn, err := net.ListenPacket("udp", ":0"); err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}
