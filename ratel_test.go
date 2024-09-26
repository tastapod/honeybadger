package honeybadger_test

import (
	"fmt"
	"strings"
	"testing"

	yaml3 "gopkg.in/yaml.v3"
)

func TestDecodesStringFromNode(t *testing.T) {
	string_node := "hello world"

	reader := strings.NewReader(string_node)
	decoder := yaml3.NewDecoder(reader)
	node := yaml3.Node{}
	decoder.Decode(&node)
	println(len(node.Content))
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", *node.Content[0])
	// a = any.Any()
}
