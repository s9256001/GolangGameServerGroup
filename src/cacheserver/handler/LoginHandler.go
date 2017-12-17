package handler

import (
	"encoding/json"

	"github.com/satori/go.uuid"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
)

// LoginHandler handles the registration of the player's account
type LoginHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *LoginHandler) Code() int {
	return gamedefine.Login
}

// OnHandle is called when Handle()
func (h *LoginHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()

	response := gamedefine.NewLoginResultPacket()

	packet := &gamedefine.LoginPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	response.PeerID = packet.PeerID
	response.RegionAddress = "ws://127.0.0.1:7772/region"
	response.PlayerKey = uuid.NewV4().String()
	response.Result = sysdefine.OK
	h.Server.SendPacket(peer, response)
	return true
}

// NewLoginHandler is a constructor of LoginHandler
func NewLoginHandler(server ginterface.IGameServer) *LoginHandler {
	ret := &LoginHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
