package server

import (
	"golang.org/x/net/websocket"

	"../handler"

	"../../base/ginterface"
	"../../servercommon"
	"../../servercommon/sysdefine"
	"../peer"
)

// DerivedGameServer implements the login server instance
type DerivedGameServer struct {
	*servercommon.SubServerBase // base class
}

// OnStart is called before Start()
func (s *DerivedGameServer) OnStart() {
	s.GetLogger().Debug("OnStart: serverName = %s\n", s.Setting.ServerName)
}

// OnStopped is called at the end of Stop()
func (s *DerivedGameServer) OnStopped() {
	s.GetLogger().Debug("OnStopped: serverName = %s\n", s.Setting.ServerName)
}

// OnCreatePeer is called to create the custom peer
func (s *DerivedGameServer) OnCreatePeer(conn *websocket.Conn) ginterface.IGamePeer {
	return peer.NewDerivedGamePeer(s, conn)
}

// OnDefaultHandle is called when there is no corresponding packet handler
func (s *DerivedGameServer) OnDefaultHandle(peer ginterface.IGamePeer, info string) {
	s.GetLogger().Debug("OnDefaultHandle: info = %s\n", info)
}

// NewDerivedGameServer is a constructor of DerivedGameServer
func NewDerivedGameServer() *DerivedGameServer {
	log := servercommon.NewConsoleGameLogger()
	ret := &DerivedGameServer{}
	ret.SubServerBase = servercommon.NewSubServerBase(ret, sysdefine.ServerTypeLogin, 7771, "login", log)
	ret.Setting.MasterURL = "ws://127.0.0.1:7770/cache"
	ret.RegisterHandler(handler.NewRegisterAccountHandler(ret))
	ret.RegisterHandler(handler.NewRegisterAccountResultHandler(ret))
	ret.RegisterHandler(handler.NewLoginHandler(ret))
	ret.RegisterHandler(handler.NewLoginResultHandler(ret))
	return ret
}
