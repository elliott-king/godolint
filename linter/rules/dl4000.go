package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// DL4000 MAINTAINER is deprecated
func dl4000Check(node *parser.Node, file string) (rst []string, err error) {
	for _, child := range node.Children {
		switch child.Value {
		case "maintainer":
			rst = append(rst, fmt.Sprintf("%s:%v DL4000 MAINTAINER is deprecated\n", file, child.StartLine))
		}
	}
	return rst, nil
}