package grove

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
)

func Error(key, message string) *Node {
	return &Node{
		Type:  "grove.Error",
		Key:   key,
		Props: flux.ReadStaticValue(struct{ Message string }{message}),
	}
}
