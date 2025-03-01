package grr

import (
	"fmt"
)

type Grr error

const (
	failedToGetExecutorHash = iota
	invalidValueType
	keyNotFound
	indexOutOfRange

	invalidLuaScript

	failedToWalkNode
	failedToFindNode
	failedToAddFluxExecutors
	failedToEvaluateFluxData
	failedToRenderChild
	invalidNodeKey
	invalidNodeType
	failedToGetErrorNodeMessage
	receivedErrorNode
	duplicateKey
	duplicateRole
)

var (
	// gex.

	ErrFailedToGetExecutorHash Grr = fmt.Errorf("grr%d", failedToGetExecutorHash)
	ErrInvalidValueType        Grr = fmt.Errorf("grr%d", invalidValueType)
	ErrKeyNotFound             Grr = fmt.Errorf("grr%d", keyNotFound)
	ErrIndexOutOfRange         Grr = fmt.Errorf("grr%d", indexOutOfRange)

	// guards.

	ErrInvalidLuaScript Grr = fmt.Errorf("grr%d", invalidLuaScript)

	// grove.

	ErrFailedToWalkNode            Grr = fmt.Errorf("grr%d", failedToWalkNode)
	ErrFailedToFindNode            Grr = fmt.Errorf("grr%d", failedToFindNode)
	ErrFailedToAddFluxExecutors    Grr = fmt.Errorf("grr%d", failedToAddFluxExecutors)
	ErrFailedToEvaluateFluxData    Grr = fmt.Errorf("grr%d", failedToEvaluateFluxData)
	ErrFailedToRenderChild         Grr = fmt.Errorf("grr%d", failedToRenderChild)
	ErrInvalidNodeKey              Grr = fmt.Errorf("grr%d", invalidNodeKey)
	ErrInvalidNodeType             Grr = fmt.Errorf("grr%d", invalidNodeType)
	ErrFailedToGetErrorNodeMessage Grr = fmt.Errorf("grr%d", failedToGetErrorNodeMessage)
	ErrReceivedErrorNode           Grr = fmt.Errorf("grr%d", receivedErrorNode)
	ErrDuplicateKey                Grr = fmt.Errorf("grr%d", duplicateKey)
	ErrDuplicateRole               Grr = fmt.Errorf("grr%d", duplicateRole)
)

func Wrap(grr Grr, format string, a ...any) error {
	return fmt.Errorf("[%w] "+format, append([]any{grr}, a...)...)
}
