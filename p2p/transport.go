package p2p

// Peer is a interface that represents the remote node
type Peer interface {

}

// Transport is anything that handles the comunication
// between nodes in the network. This can be of the
// form of a TCP connection, a UDP connection, etc.
type Transport interface {
	ListenAndAccept() error
}