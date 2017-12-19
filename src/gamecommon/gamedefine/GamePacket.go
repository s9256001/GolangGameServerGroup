package gamedefine

import (
	"../../base/basedefine"
	"../../servercommon/sysinfo"
)

// RegisterAccountPacket is the packet of request of registration of the player's account
type RegisterAccountPacket struct {
	basedefine.GameBasePacket        // base packet
	PeerID                    string // player's peer id
	Account                   string // player's account
	Password                  string // player's password
}

// NewRegisterAccountPacket is a constructor of RegisterAccountPacket
func NewRegisterAccountPacket() *RegisterAccountPacket {
	ret := &RegisterAccountPacket{}
	ret.Code = RegisterAccount
	return ret
}

// RegisterAccountResultPacket is the packet of response of registration of player's login
type RegisterAccountResultPacket struct {
	basedefine.GameBasePacket        // base packet
	PeerID                    string `json:",omitempty"` // player's peer id
	Result                    int    // result
}

// NewRegisterAccountResultPacket is a constructor of RegisterAccountResultPacket
func NewRegisterAccountResultPacket() *RegisterAccountResultPacket {
	ret := &RegisterAccountResultPacket{}
	ret.Code = RegisterAccountResult
	return ret
}

// LoginPacket is the packet of request of player's login
type LoginPacket struct {
	basedefine.GameBasePacket        // base packet
	PeerID                    string // player's peer id
	Account                   string // player's account
	Password                  string // player's password
}

// NewLoginPacket is a constructor of LoginPacket
func NewLoginPacket() *LoginPacket {
	ret := &LoginPacket{}
	ret.Code = Login
	return ret
}

// LoginResultPacket is the packet of response of player's login
type LoginResultPacket struct {
	basedefine.GameBasePacket        // base packet
	PeerID                    string `json:",omitempty"` // player's peer id
	RegionAddress             string // region server's address
	PlayerKey                 string // player's key for communicating to the region server
	Result                    int    // result
}

// NewLoginResultPacket is a constructor of LoginResultPacket
func NewLoginResultPacket() *LoginResultPacket {
	ret := &LoginResultPacket{}
	ret.Code = LoginResult
	return ret
}

// EnterRegionPacket is the packet of request of player's entering region
type EnterRegionPacket struct {
	basedefine.GameBasePacket        // base packet
	PeerID                    string // player's peer id
	PlayerKey                 string // player's key for communicating to the region server
}

// NewEnterRegionPacket is a constructor of EnterRegionPacket
func NewEnterRegionPacket() *EnterRegionPacket {
	ret := &EnterRegionPacket{}
	ret.Code = EnterRegion
	return ret
}

// EnterRegionResultPacket is the packet of response of player's entering region
type EnterRegionResultPacket struct {
	basedefine.GameBasePacket                        // base packet
	PeerID                    string                 `json:",omitempty"` // player's peer id
	PlayerInfo                sysinfo.PlayerInfoBase // the information of the player
	Result                    int                    // result
}

// NewEnterRegionResultPacket is a constructor of EnterRegionResultPacket
func NewEnterRegionResultPacket() *EnterRegionResultPacket {
	ret := &EnterRegionResultPacket{}
	ret.Code = EnterRegionResult
	return ret
}

// EnterGamePacket is the packet of request of player's entering game
type EnterGamePacket struct {
	basedefine.GameBasePacket     // base packet
	GameID                    int // game id
}

// NewEnterGamePacket is a constructor of EnterGamePacket
func NewEnterGamePacket() *EnterGamePacket {
	ret := &EnterGamePacket{}
	ret.Code = EnterGame
	return ret
}

// EnterGameResultPacket is the packet of response of player's entering game
type EnterGameResultPacket struct {
	basedefine.GameBasePacket     // base packet
	GameID                    int // game id
	Result                    int // result
}

// NewEnterGameResultPacket is a constructor of EnterGameResultPacket
func NewEnterGameResultPacket() *EnterGameResultPacket {
	ret := &EnterGameResultPacket{}
	ret.Code = EnterGameResult
	return ret
}

// LeaveGamePacket is the packet of request of player's leaving game
type LeaveGamePacket struct {
	basedefine.GameBasePacket     // base packet
	GameID                    int // game id
}

// NewLeaveGamePacket is a constructor of LeaveGamePacket
func NewLeaveGamePacket() *LeaveGamePacket {
	ret := &LeaveGamePacket{}
	ret.Code = LeaveGame
	return ret
}

// LeaveGameResultPacket is the packet of response of player's leaving game
type LeaveGameResultPacket struct {
	basedefine.GameBasePacket     // base packet
	GameID                    int // game id
	Result                    int // result
}

// NewLeaveGameResultPacket is a constructor of LeaveGameResultPacket
func NewLeaveGameResultPacket() *LeaveGameResultPacket {
	ret := &LeaveGameResultPacket{}
	ret.Code = LeaveGameResult
	return ret
}
