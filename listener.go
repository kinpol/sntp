package sntp

import "net"

type UdpClient interface {
	DatagramReceived(data []byte, addr net.Addr)
	SetUdpTransport(Transport)
}

type UnixHandler interface {
	UnixReceived(data []byte, conn *net.UnixConn)
}

type TcpClient interface {
	DataReceived(data []byte, conn *net.TCPConn)
	SetTcpTransport(Transport)
}
