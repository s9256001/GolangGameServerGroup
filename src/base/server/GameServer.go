package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"

	"../basedefine"
	"../ginterface"
	"../module"
	"../node"
)

// GameServer is an abstract class of the game server
// It fits the interface of IGameServer, and depends on the derived type to implement the interface of IGameServerHook
type GameServer struct {
	ginterface.IGameServerHook // hook
	*node.Node                 // base class

	Handle   *http.Server                       // http server handle
	Setting  module.ServerSetting               // server setting
	Handlers map[int]ginterface.IGameHandler    // map of packet handlers with key denoting the packet command
	Peers    map[uuid.UUID]ginterface.IGamePeer // map of game peers with key denoting the peer id

	MasterPeer ginterface.IGamePeer // game peer of the master server
}

// GetModule returns the specific module to resolve import cycle
func (s *GameServer) GetModule(m interface{}) interface{} {
	switch m.(type) {
	case module.ServerSetting:
		return s.Setting
	}
	return nil
}

// RegisterHandler registers the packet handler
func (s *GameServer) RegisterHandler(handler ginterface.IGameHandler) {
	s.Handlers[handler.Code()] = handler
}

// Start starts the server
func (s *GameServer) Start() bool {
	addr := ":" + strconv.Itoa(s.Setting.Port)
	s.Handle = &http.Server{Addr: addr}

	// router
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/"+s.Setting.ServerName, websocket.Handler(s.PeerHandler))

	// connect to the master server
	if s.Setting.MasterURL != "" {
		ws, err := websocket.Dial(s.Setting.MasterURL, "", "http://localhost/")
		if err != nil {
			s.Log.Error("Start: cannot connect to the master! err = %v\n", err)
			return false
		}

		peer := s.OnCreatePeer(ws)
		s.MasterPeer = peer
		peer.OnConnected()

		// service loop to handle the request from the master server
		go func(peer ginterface.IGamePeer) {
			defer ws.Close()

			s.PeerReceiveLoop(peer)
		}(peer)

		s.OnRegisterToMaster()
	}

	s.OnStart()

	// service loop
	if err := http.ListenAndServe(addr, nil); err != nil {
		return false
	}
	return true
}

// Stop stops the server
func (s *GameServer) Stop() bool {
	if err := s.Handle.Shutdown(nil); err != nil {
		s.Log.Error("Stop: shutdown failed! err = %v\n", err)
		return false
	}
	s.OnStopped()
	return true
}

// GetMasterPeer returns the game peer of the master server
func (s *GameServer) GetMasterPeer() ginterface.IGamePeer {
	return s.MasterPeer
}

// GetPeer returns the game peer of the peerID
func (s *GameServer) GetPeer(peerID uuid.UUID) ginterface.IGamePeer {
	peer, ok := s.Peers[peerID]
	if !ok {
		return nil
	}
	return peer
}

// SendPacket sends packet to the connection
func (s *GameServer) SendPacket(peer ginterface.IGamePeer, packet interface{}) bool {
	if peer == nil {
		s.Log.Error("SendPacket: peer is nil")
		return false
	}
	message, err := json.Marshal(packet)
	if err != nil {
		s.Log.Error("SendPacket: fail to serialize the packet! err = %v\n", err)
		return false
	}
	err = websocket.Message.Send(peer.GetConn(), string(message))
	if err != nil {
		s.Log.Error("SendPacket: fail to send the packet! err = %v\n", err)
		return false
	}
	return true
}

// PeerHandler is the handler of connection
func (s *GameServer) PeerHandler(ws *websocket.Conn) {
	defer ws.Close()

	// create peer
	peer := s.OnCreatePeer(ws)
	s.Peers[peer.GetPeerID()] = peer
	peer.OnConnected()

	s.PeerReceiveLoop(peer)
}

// PeerReceiveLoop is the service loop for waitting to receive a message
func (s *GameServer) PeerReceiveLoop(peer ginterface.IGamePeer) {
	var receivedMessage string
	for {
		// wait to receive a message
		if err := websocket.Message.Receive(peer.GetConn(), &receivedMessage); err != nil {
			s.Log.Error("PeerHandler: receive failed! err = %v\n", err)
			peer.OnDisconnected()
			break
		}

		// deserialize the packet
		basePacket := basedefine.GameBasePacket{}
		if err := json.Unmarshal([]byte(receivedMessage), &basePacket); err != nil {
			s.Log.Error("PeerHandler: deserialized failed! err = %v\n", err)
			peer.OnDisconnected()
			break
		}

		// dispatch the packet to handlers
		if handler, ok := s.Handlers[basePacket.Code]; ok {
			handler.Handle(peer, string(receivedMessage))
		} else {
			s.OnDefaultHandle(peer, string(receivedMessage))
		}
	}
}

// NewGameServer is a constructor of GameServer
func NewGameServer(hook ginterface.IGameServerHook, serverType int, port int, serverName string, log ginterface.IGameLogger) *GameServer {
	ret := &GameServer{
		IGameServerHook: hook,
		Node:            node.NewNode(log),

		Setting: module.ServerSetting{
			ServerType: serverType,
			Port:       port,
			ServerName: serverName,
		},
		Handlers: make(map[int]ginterface.IGameHandler),
		Peers:    make(map[uuid.UUID]ginterface.IGamePeer),
	}
	return ret
}
