package grove

import "github.com/kmirzavaziri/grove/go/pkg/gex"

type Action struct {
	Type    string
	Payload *gex.Value
}

type ActionRenderErrorMessagesPayload struct {
	ErrorMessage  string
	ErrorMessages map[string]string
	ErrorInfo     string
}

func ActionRenderErrorMessages(payload ActionRenderErrorMessagesPayload) *Action {
	return &Action{
		Type:    "grove.ActionRenderErrorMessages",
		Payload: gex.Marshal(payload),
	}
}
