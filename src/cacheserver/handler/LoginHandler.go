package handler

import (
	"encoding/json"
	"fmt"

	"github.com/satori/go.uuid"

	"../../base/ginterface"
	"../../base/handler"
	"../../gamecommon/gamedefine"
	"../../servercommon/sysdefine"
	"../../servercommon/sysinfo"
)

// LoginHandler handles player's login
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
	subServers := h.Server.GetModule(map[uuid.UUID]sysinfo.SubServerInfo{}).(map[uuid.UUID]sysinfo.SubServerInfo)
	players := h.Server.GetModule(map[uuid.UUID]sysinfo.PlayerInfo{}).(map[uuid.UUID]sysinfo.PlayerInfo)

	response := gamedefine.NewLoginResultPacket()

	packet := &gamedefine.LoginPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("LoginHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}

	playerKey := uuid.NewV4()
	players[playerKey] = sysinfo.PlayerInfo{
		PlayerInfoBase: sysinfo.PlayerInfoBase{
			PlayerKey: playerKey,
			Name:      packet.Account,
			Gold:      1000000,
		},
	}

	var regionAddress string
	for _, subServerInfo := range subServers {
		if subServerInfo.ServerType == sysdefine.ServerTypeRegion {
			regionAddress = fmt.Sprintf("ws://%s:%d/%s", subServerInfo.Address, subServerInfo.Port, subServerInfo.ServerName)
		}
	}

	response.PeerID = packet.PeerID
	response.RegionAddress = regionAddress
	response.PlayerKey = playerKey.String()
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
