package dot

import (
	"fmt"
	"strings"

	"github.com/m8u/gorgonia"
	internalEncoding "github.com/m8u/gorgonia/internal/encoding"
	"gonum.org/v1/gonum/graph/encoding"
)

type node struct {
	n *gorgonia.Node
}

func (n *node) ID() int64 {
	return n.n.ID()
}

// DOTID is used for the graphviz output. It fulfils the gonum encoding interface
func (n *node) DOTID() string {
	return fmt.Sprintf("Node_%p", n.n)
}

// Attributes is for graphviz output. It specifies the "label" of the node (a table)
func (n *node) Attributes() []encoding.Attribute {
	var htmlEscaper = strings.NewReplacer(
		`&`, "&amp;",
		`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
		`<`, "&lt;",
		`>`, "&gt;",
		`{`, "&#123;",
		`}`, "&#125;",
		`"`, "&#34;", // "&#34;" is shorter than "&quot;".
		`const`, "const|", // "&#34;" is shorter than "&quot;".
	)
	attrs := []encoding.Attribute{
		{
			Key:   "id",
			Value: fmt.Sprintf(`"%p"`, n.n),
		},
		{
			Key:   "shape",
			Value: "Mrecord",
		},
		{
			Key: "label",
			Value: fmt.Sprintf(`"{{%s|%#x}|{Op|%s}|{Shape|%v}}"`,
				htmlEscaper.Replace(n.n.Name()),
				n.ID(),
				htmlEscaper.Replace(fmt.Sprintf("%s", n.n.Op())),
				n.n.Shape()),
		},
	}
	return attrs
}

func (n *node) Groups() internalEncoding.Groups {
	return n.n.Groups()
}
