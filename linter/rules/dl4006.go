package rules

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"strings"
)

// validateDL4006 Set the `SHELL` option -o pipefail before `RUN` with a pipe in it
func validateDL4006(node *parser.Node, file string) (rst []string, err error) {
	isShellPipeFail := false
	for _, child := range node.Children {
		switch child.Value {
		case "shell":
			isShellPipeFail = true
		case "run":
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "|":
					if !isShellPipeFail {
						rst = append(rst, fmt.Sprintf("%s:%v DL4006 Set the `SHELL` option -o pipefail before `RUN` with a pipe in it\n", file, child.StartLine))
					}
				}
			}
		}
	}
	return rst, nil
}