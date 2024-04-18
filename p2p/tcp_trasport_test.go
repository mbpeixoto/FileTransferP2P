package p2p

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	// Create a new TCPTransport
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

	// Server
	err := tr.ListenAndAccept()
	assert.Nil(t,err)

}