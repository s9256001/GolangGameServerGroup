package sysdefine

const (
	// RegisterSubServer is the packet command of request of registration to the master server
	RegisterSubServer = iota + 1
	// RegisterSubServerResult is the packet command of response of registration from the master server
	RegisterSubServerResult
)
