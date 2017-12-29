package handler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// LoginHandler handles the request of player's login
type LoginHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginHandler) Code() int {
	return gamedefine.Login
}

// OnHandle is called when Handling the packet
func (h *LoginHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Node.GetLogger()

	response := gamedefine.NewLoginResultPacket()
	response.Result = sysdefine.Failed

	packet := &gamedefine.LoginPacket{}

	defer func() {
		if response.Result == sysdefine.OK {
			h.Node.(ginterface.IGameServer).SendPacket(h.Node.(ginterface.IGameServer).GetMasterPeer(), packet)
		} else {
			h.Node.(ginterface.IGameServer).SendPacket(peer, response)
		}
	}()

	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	packet.PeerID = peer.GetPeerID().String()
	response.Result = sysdefine.OK
	return true
}

// NewLoginHandler is a constructor of LoginHandler
func NewLoginHandler(node ginterface.INode) *LoginHandler {
	ret := &LoginHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, node)
	return ret
}
