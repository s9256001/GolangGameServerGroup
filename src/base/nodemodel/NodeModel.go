package nodemodel

// NodeModel is an abstract class of the node model
// It fits the interface of INodeModel
type NodeModel struct {
}

// NewNodeModel is a constructor of Node
func NewNodeModel() *NodeModel {
	ret := &NodeModel{}
	return ret
}
