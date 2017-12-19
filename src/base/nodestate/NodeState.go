package nodestate

import (
	"../ginterface"
)

// NodeState is an abstract class of the node state
// It fits the interface of INodeState
type NodeState struct {
	Node ginterface.INode // node
}

// NewNodeState is a constructor of NodeState
func NewNodeState(node ginterface.INode) *NodeState {
	ret := &NodeState{
		Node: node,
	}
	return ret
}
