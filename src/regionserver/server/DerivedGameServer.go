package server

import (
	"golang.org/x/net/websocket"

	"../../base/ginterface"
	"../../servercommon"
	"../../servercommon/sysdefine"
	"../handler"
	"../peer"
)

// DerivedGameServer implements the region server instance
type DerivedGameServer struct {
	*servercommon.SubServerBase // base class
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
	ret.SubServerBase = servercommon.NewSubServerBase(ret, sysdefine.ServerTypeRegion, 7772, "region", log)
	ret.Setting.MasterURL = "ws://127.0.0.1:7770/cache"
	ret.RegisterHandler(handler.NewEnterRegionHandler(ret))
	ret.RegisterHandler(handler.NewEnterRegionResultHandler(ret))
	return ret
}
