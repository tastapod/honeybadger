package yaml_test

import (
	"fmt"
	"strings"
	"testing"

	yaml3 "gopkg.in/yaml.v3"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, [4, true, 077, 0o77, 17.34, 314e-2]]
  infinity: .inf
  hex: 0xff
  nope: .nan
some-time: 2001-12-15T02:59:43.1Z
space-time: 2001-12-15 2:59:43.10
just-date: 2002-12-14
`

func TestParsesYaml(t *testing.T) {
	reader := strings.NewReader(data)
	decoder := yaml3.NewDecoder(reader)
	node := yaml3.Node{}
	decoder.Decode(&node)
	for _, n := range node.Content {
		printNode(n, 0)
	}
}

type target struct {
	Name string
}

var kinds = map[yaml3.Kind]string{
	yaml3.DocumentNode: "Document",
	yaml3.SequenceNode: "Sequence",
	yaml3.MappingNode:  "Mapping",
	yaml3.ScalarNode:   "Scalar",
	yaml3.AliasNode:    "Alias",
}

func printNode(node *yaml3.Node, level int) {
	print(strings.Repeat(" ", level*2))
	fmt.Printf("Node{Kind: %s, Tag: %s, Value: '%v'}\n", kinds[node.Kind], node.Tag, node.Value)
	for _, child := range node.Content {
		printNode(child, level+1)
	}
}
