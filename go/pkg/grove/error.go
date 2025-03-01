package grove

import "github.com/kmirzavaziri/grove/go/pkg/gex"

func Error(key, message string) *Node {
	return &Node{
		Type: "grove.Error",
		Key:  key,
		Data: gex.StaticV(struct{ Message string }{message}),
	}
}
