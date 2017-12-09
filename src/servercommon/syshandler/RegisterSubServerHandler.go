package syshandler

import (
	"encoding/json"

	"../../base/ginterface"
	"../../base/handler"
	"../../base/module"
	"../sysdefine"
)

// RegisterSubServerHandler handles the registration of subserver
type RegisterSubServerHandler struct {
	*handler.GameHandler // base class
}

// Code is the associated packet command
func (h *RegisterSubServerHandler) Code() int {
	return sysdefine.RegisterSubServer
}

// OnHandle is called when Handle()
func (h *RegisterSubServerHandler) OnHandle(peer ginterface.IGamePeer, info string) bool {
	log := h.Server.GetLogger()
	setting := h.Server.GetModule(module.ServerSetting{}).(module.ServerSetting)

	response := sysdefine.NewRegisterSubServerResultPacket()
	response.ServerType = 0
	response.Result = sysdefine.Failed

	packet := &sysdefine.RegisterSubServerPacket{}
	if err := json.Unmarshal([]byte(info), &packet); err != nil {
		log.Error("RegisterSubServerHandler.OnHandle(): failed to deserialize! info = %s\n", info)
		return false
	}
	log.Debug("RegisterSubServerHandler.OnHandle(): code = %d, serverType = %d\n", h.Code(), packet.ServerType)

	response.ServerType = setting.ServerType
	response.Result = sysdefine.OK
	h.Server.SendPacket(peer, response)
	return true
}

// NewRegisterSubServerHandler is a constructor of RegisterSubServerHandler
func NewRegisterSubServerHandler(server ginterface.IGameServer) *RegisterSubServerHandler {
	ret := &RegisterSubServerHandler{}
	ret.GameHandler = handler.NewGameHandler(ret, server)
	return ret
}
