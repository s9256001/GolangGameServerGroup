package model

import (
	"../../base/nodemodel"
)

// DerivedModel is the custom model of this game
type DerivedModel struct {
	*nodemodel.NodeModel // base class
}

// NewDerivedModel is a constructor of DerivedModel
func NewDerivedModel() *DerivedModel {
	ret := &DerivedModel{}
	ret.NodeModel = nodemodel.NewNodeModel()
	return ret
}
