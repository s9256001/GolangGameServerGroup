package peer

import (
	"golang.org/x/net/websocket"

	"../../base/ginterface"
	"../../base/peer"
)

// DerivedGamePeer is the custom game peer of this server
type DerivedGamePeer struct {
	*peer.GamePeer // base class
}

// OnConnected is called when connected
func (p *DerivedGamePeer) OnConnected() {
	log := p.Server.GetLogger()
	log.Debug("DerivedGamePeer.OnConnected: peerID = %s\n", p.PeerID)
}

// OnDisconnected is called when disconnected
func (p *DerivedGamePeer) OnDisconnected() {
	log := p.Server.GetLogger()
	log.Debug("DerivedGamePeer.OnDisconnected: peerID = %s\n", p.PeerID)
}

// NewDerivedGamePeer is a constructor of DerivedGamePeer
func NewDerivedGamePeer(server ginterface.IGameServer, conn *websocket.Conn) *DerivedGamePeer {
	ret := &DerivedGamePeer{}
	ret.GamePeer = peer.NewGamePeer(server, conn)
	return ret
}
