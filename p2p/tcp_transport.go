package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer representa o peer remoto que se conecta via TCP
type TCPPeer struct {
	// conn é a conexão do peer
	conn net.Conn

	// se foi feita uma conexão, outbound = true
	// se foi recebida uma conexão, outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	ln, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener= ln
	return nil
}


func (t *TCPTransport) startAcceptLoop() {
for {
	conn, err := t.listener.Accept()
	if err != nil {
		fmt.Printf("Erro na conexão TCP: %v\n", err)
	}

	go t.handleConn(conn)
}
}

func (t *TCPTransport) handleConn(conn net.Conn) error {
	peer := NewTCPPeer(conn, true)
	
	fmt.Printf("Nova conexão %v\n", conn)
	
	return nil
}