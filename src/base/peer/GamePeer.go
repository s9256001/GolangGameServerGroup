package peer

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"

	"../ginterface"
)

// GamePeer is an abstract class of the game peer
// It fits the interface of IGamePeer
type GamePeer struct {
	PeerID uuid.UUID // the id of the peer
	// todo
	IP     string                 // ip of this peer
	Port   int                    // port of this peer
	Server ginterface.IGameServer // server
	Conn   *websocket.Conn        // connection
}

// GetPeerID returns the id of the peer
func (p *GamePeer) GetPeerID() uuid.UUID {
	return p.PeerID
}

// GetConn gets the connection of the game peer
func (p *GamePeer) GetConn() *websocket.Conn {
	return p.Conn
}

// NewGamePeer is a constructor of GamePeer
func NewGamePeer(server ginterface.IGameServer, conn *websocket.Conn) *GamePeer {
	ret := &GamePeer{
		PeerID: uuid.NewV4(),
		Server: server,
		Conn:   conn,
	}
	return ret
}
