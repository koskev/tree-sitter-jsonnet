package tree_sitter_jsonnet_test

import (
	"testing"

	tree_sitter_jsonnet "github.com/koskev/tree-sitter-jsonnet/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_jsonnet.Language())
	if language == nil {
		t.Errorf("Error loading Jsonnet grammar")
	}
}
