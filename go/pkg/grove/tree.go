package grove

import (
	"strings"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Node struct {
	Type     string
	Key      string
	Role     string
	Data     gex.Flux[*gex.Value]
	Children []*Node
}

type walkedNode struct {
	type_        string
	key          string
	role         string
	data         gex.Flux[*gex.Value]
	children     []*walkedNode
	path         []string
	flatChildren []*walkedNode
}

func newWalkedNode(node *Node, parentPath []string) (*walkedNode, error) {
	if node.Key == "" || strings.Contains(node.Key, "/") {
		return nil, grr.Wrap(grr.ErrInvalidNodeKey, "invalid node key %q", node.Key)
	}

	if node.Type == "" {
		return nil, grr.Wrap(grr.ErrInvalidNodeType, "invalid node type %q", node.Type)
	}

	if node.Type == "grove.Error" {
		nodeData, err := node.Data.Data(gex.Nil(), nil)
		if err != nil {
			return nil, grr.Wrap(grr.ErrFailedToGetErrorNodeMessage, "received error node but failed to get data")
		}

		message, err := nodeData.Chain().Get("message").String()
		if err != nil {
			return nil, grr.Wrap(
				grr.ErrFailedToGetErrorNodeMessage, "received error node but failed to get message from data")
		}

		return nil, grr.Wrap(grr.ErrReceivedErrorNode, "received error: %q", message)
	}

	keySet := make(map[string]struct{})

	for _, child := range node.Children {
		if _, exists := keySet[child.Key]; exists {
			return nil, grr.Wrap(grr.ErrDuplicateKey, "duplicate key %q in children", child.Key)
		}

		keySet[child.Key] = struct{}{}
	}

	roleSet := make(map[string]struct{})

	for _, child := range node.Children {
		if child.Role != "" {
			if _, exists := roleSet[child.Role]; exists {
				return nil, grr.Wrap(grr.ErrDuplicateRole, "duplicate role %q in children", child.Role)
			}

			roleSet[child.Role] = struct{}{}
		}
	}

	path := make([]string, len(parentPath)+1)
	for i := range parentPath {
		path[i] = parentPath[i]
	}

	path[len(parentPath)] = node.Key

	wNode := &walkedNode{
		type_:        node.Type,
		key:          node.Key,
		role:         node.Role,
		data:         node.Data,
		children:     make([]*walkedNode, len(node.Children)),
		path:         path,
		flatChildren: []*walkedNode{},
	}

	for i, child := range node.Children {
		wChild, err := newWalkedNode(child, path)
		if err != nil {
			return nil, grr.Wrap(grr.ErrFailedToWalkNode, "failed to walk over node %q: %w", child.Key, err)
		}

		wNode.children[i] = wChild

		wNode.flatChildren = append(wNode.flatChildren, wChild)
		wNode.flatChildren = append(wNode.flatChildren, wChild.flatChildren...)
	}

	return wNode, nil
}
