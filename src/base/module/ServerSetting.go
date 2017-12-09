package module

// ServerSetting is a module consisting of settings of the server
type ServerSetting struct {
	ServerType int    // server type
	Port       int    // server listening port
	ServerName string // server name
	MasterURL  string // url of the master server
}
