package main

import (
	"fmt"
	// "log"
	"strings"

	"gopkg.in/yaml.v3"
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

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	reader := strings.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	node := yaml.Node{}
	decoder.Decode(&node)
	for _, n := range node.Content {
		printNode(n, 0)
	}
}

var kinds = map[yaml.Kind]string{
	yaml.DocumentNode: "Document",
	yaml.SequenceNode: "Sequence",
	yaml.MappingNode:  "Mapping",
	yaml.ScalarNode:   "Scalar",
	yaml.AliasNode:    "Alias",
}

func printNode(node *yaml.Node, level int) {
	print(strings.Repeat(" ", level*2))
	fmt.Printf("Node{Kind: %s, Tag: %s, Value: '%v'}\n", kinds[node.Kind], node.Tag, node.Value)
	for _, child := range node.Content {
		printNode(child, level+1)
	}
}
