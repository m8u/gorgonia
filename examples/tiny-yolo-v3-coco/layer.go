package main

import "github.com/m8u/gorgonia"

type layerN interface {
	String() string
	Type() string
	ToNode(g *gorgonia.ExprGraph, input ...*gorgonia.Node) (*gorgonia.Node, error)
}
