package grove

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"strings"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Node struct {
	Type     string
	Key      string
	Role     string
	Props    flux.Read[*gex.Value]
	Submit   flux.Act[*Action]
	Input    *Input
	Children []*Node

	parent                *Node
	path                  string
	flatChildrenByPath    map[string]*Node
	hasSubmittableContext bool
	flatInputsByKey       map[string]*Input
}

func (n *Node) validateAndPopulate() error {
	err := n.validate()
	if err != nil {
		return err
	}

	n.populateParents(nil)

	err = n.PreOrderDFS(populatePath)
	if err != nil {
		return err
	}

	err = n.PostOrderDFS(populateFlatChildren)
	if err != nil {
		return err
	}

	err = n.PreOrderDFS(populateHasSubmittableContext)
	if err != nil {
		return err
	}

	err = n.PostOrderDFS(populateFlatInputs)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) PreOrderDFS(visit func(node *Node) error) error {
	err := visit(n)
	if err != nil {
		return err
	}

	for _, c := range n.Children {
		err := c.PreOrderDFS(visit)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) PostOrderDFS(visit func(node *Node) error) error {
	for _, c := range n.Children {
		err := c.PostOrderDFS(visit)
		if err != nil {
			return err
		}
	}

	err := visit(n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) validate() error {
	if n.Key == "" || strings.Contains(n.Key, "/") {
		return grr.Wrap(grr.ErrInvalidNodeKey, "invalid node key %q", n.Key)
	}

	if n.Type == "" {
		return grr.Wrap(grr.ErrInvalidNodeType, "invalid node type %q", n.Type)
	}

	if n.Input != nil && n.Key == "" {
		return grr.Wrap(grr.ErrInvalidInputKey, "invalid input key %q", n.Input.Key)
	}

	if n.Type == "grove.Error" {
		props, err := n.Props.Data(greq.Request{}, nil)
		if err != nil {
			return grr.Wrap(grr.ErrFailedToGetErrorNodeMessage, "received error node but failed to get props data")
		}

		message, err := props.Chain().Get("message").String()
		if err != nil {
			return grr.Wrap(grr.ErrFailedToGetErrorNodeMessage, "received error node but failed to get message")
		}

		return grr.Wrap(grr.ErrReceivedErrorNode, "received error: %q", message)
	}

	keySet := make(map[string]struct{})
	for _, child := range n.Children {
		if _, exists := keySet[child.Key]; exists {
			return grr.Wrap(grr.ErrDuplicateNodeKey, "duplicate key %q in children", child.Key)
		}

		keySet[child.Key] = struct{}{}
	}

	roleSet := make(map[string]struct{})
	for _, child := range n.Children {
		if child.Role != "" {
			if _, exists := roleSet[child.Role]; exists {
				return grr.Wrap(grr.ErrDuplicateNodeRole, "duplicate role %q in children", child.Role)
			}

			roleSet[child.Role] = struct{}{}
		}
	}

	return nil
}

func (n *Node) populateParents(parent *Node) {
	n.parent = parent

	for _, c := range n.Children {
		c.populateParents(n)
	}
}

func populatePath(n *Node) error {
	parentPath := ""

	if n.parent != nil {
		parentPath = n.parent.path + "/"
	}

	n.path = parentPath + n.Key

	return nil
}

func populateFlatChildren(n *Node) error {
	n.flatChildrenByPath = map[string]*Node{}

	for _, c := range n.Children {
		n.flatChildrenByPath[c.path] = c
	}

	for _, c := range n.Children {
		for k, v := range c.flatChildrenByPath {
			n.flatChildrenByPath[k] = v
		}
	}

	return nil
}

func populateHasSubmittableContext(n *Node) error {
	n.hasSubmittableContext = (n.parent != nil && n.parent.hasSubmittableContext) || n.Submit != nil

	return nil
}

func populateFlatInputs(n *Node) error {
	if !n.hasSubmittableContext {
		return nil
	}

	n.flatInputsByKey = map[string]*Input{}

	if n.Input != nil {
		n.flatInputsByKey[n.Input.Key] = n.Input
	}

	for _, c := range n.Children {
		for k, v := range c.flatInputsByKey {
			if _, exists := n.flatInputsByKey[k]; exists {
				return grr.Wrap(grr.ErrDuplicateInputKey, "duplicate input key %q in node %q", k, n.path)
			}

			n.flatInputsByKey[k] = v
		}
	}

	return nil
}
