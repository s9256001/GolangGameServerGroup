package sysdefine

import (
	"../../base/basedefine"
)

// RegisterSubServerPacket is the packet of request of registration to the master server
type RegisterSubServerPacket struct {
	basedefine.GameBasePacket        // base packet
	ServerType                int    // server type
	Port                      int    // server listening port
	ServerName                string // server name
}

// NewRegisterSubServerPacket is a constructor of RegisterSubServerPacket
func NewRegisterSubServerPacket() *RegisterSubServerPacket {
	ret := &RegisterSubServerPacket{}
	ret.Code = RegisterSubServer
	return ret
}

// RegisterSubServerResultPacket is the packet of response of registration from the master server
type RegisterSubServerResultPacket struct {
	basedefine.GameBasePacket     // base packet
	Result                    int // result
}

// NewRegisterSubServerResultPacket is a constructor of RegisterSubServerResultPacket
func NewRegisterSubServerResultPacket() *RegisterSubServerResultPacket {
	ret := &RegisterSubServerResultPacket{}
	ret.Code = RegisterSubServerResult
	return ret
}
