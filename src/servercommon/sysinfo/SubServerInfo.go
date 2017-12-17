package sysinfo

import (
	"github.com/satori/go.uuid"
)

// SubServerInfo is the basic information of the subserver
type SubServerInfoBase struct {
	PeerID     uuid.UUID // the id of the peer
	ServerType int       // server type
	Address    string    // server address
	Port       int       // server listening port
	ServerName string    // server name
}

// SubServerInfo is the information of the subserver
type SubServerInfo struct {
	SubServerInfoBase // base class

	Data interface{} // custom data
}
