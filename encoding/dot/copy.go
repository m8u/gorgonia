package dot

import (
	"github.com/m8u/gorgonia"
	"gonum.org/v1/gonum/graph"
)

func copyGraph(dst graph.Builder, src graph.Graph) {
	nodes := src.Nodes()
	for nodes.Next() {
		current := nodes.Node().(*gorgonia.Node)
		dst.AddNode(&node{
			n: current,
		})
	}
	nodes.Reset()
	for nodes.Next() {
		u := &node{
			n: nodes.Node().(*gorgonia.Node),
		}
		uid := u.ID()
		to := src.From(uid)
		for to.Next() {
			v := &node{
				n: to.Node().(*gorgonia.Node),
			}
			dst.SetEdge(dst.NewEdge(u, v))
			//dst.SetEdge(src.Edge(uid, v.ID()))
		}
	}
}
