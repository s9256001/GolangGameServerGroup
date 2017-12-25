package node

import (
	"reflect"

	"../ginterface"
)

// Node is an abstract class of the node
// It fits the interface of INode
type Node struct {
	Log   ginterface.IGameLogger // game logger
	Model ginterface.INodeModel  // node model
	State ginterface.INodeState  // node state
}

// GetLogger returns the game logger
func (n *Node) GetLogger() ginterface.IGameLogger {
	return n.Log
}

// GetModule returns the specific module to resolve import cycle
func (n *Node) GetModule(module interface{}) interface{} {
	return nil
}

// GetModel gets the model
func (n *Node) GetModel() ginterface.INodeModel {
	return n.Model
}

// GetState gets current state
func (n *Node) GetState() ginterface.INodeState {
	return n.State
}

// ChangeState changes the state of the node
func (n *Node) ChangeState(newState ginterface.INodeState) {
	if newState == nil {
		n.Log.Error("ChangeState: newState == nil!\n")
		return
	}
	preState := n.State
	n.State = newState

	n.Log.Debug("ChangeState: {%s} -> {%s\n",
		reflect.TypeOf(preState).Name(), reflect.TypeOf(newState).Name())

	if preState != nil {
		preState.OnExit()
	}
	newState.OnEnter()
}

// OnUpdate is called when updating this node
func (n *Node) OnUpdate(gapSecs int) {
	n.State.OnUpdate(gapSecs)
}

// NewNode is a constructor of Node
func NewNode(log ginterface.IGameLogger) *Node {
	ret := &Node{Log: log}
	return ret
}
