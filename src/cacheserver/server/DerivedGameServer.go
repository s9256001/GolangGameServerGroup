package server

import (
	"golang.org/x/net/websocket"

	"../../base/ginterface"
	"../../servercommon"
	"../peer"
)

// DerivedGameServer implements the cache server instance
type DerivedGameServer struct {
	*servercommon.MasterServerBase // base class
}

// OnStart is called before Start()
func (s *DerivedGameServer) OnStart() {
	s.Log.Debug("OnStart: serverName = %s\n", s.Setting.ServerName)
}

// OnStopped is called at the end of Stop()
func (s *DerivedGameServer) OnStopped() {
	s.Log.Debug("OnStopped: serverName = %s\n", s.Setting.ServerName)
}

// OnCreatePeer is called to create the custom peer
func (s *DerivedGameServer) OnCreatePeer(conn *websocket.Conn) ginterface.IGamePeer {
	return peer.NewDerivedGamePeer(s, conn)
}

// OnDefaultHandle is called when there is no corresponding packet handler
func (s *DerivedGameServer) OnDefaultHandle(peer ginterface.IGamePeer, info string) {
	s.Log.Debug("OnDefaultHandle: info = %s\n", info)
}

// NewDerivedGameServer is a constructor of DerivedGameServer
func NewDerivedGameServer() *DerivedGameServer {
	log := servercommon.NewConsoleGameLogger()
	ret := &DerivedGameServer{}
	ret.MasterServerBase = servercommon.NewMasterServerBase(ret, 0, 7770, "cache", log)
	return ret
}
