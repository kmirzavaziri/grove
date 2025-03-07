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
	failedToEvaluateFluxData

	invalidLuaScript

	failedToValidateAndPopulateGrove
	failedToFindNode
	failedToRenderChild
	invalidNodeKey
	invalidNodeType
	invalidInputKey
	failedToGetErrorNodeMessage
	receivedErrorNode
	duplicateNodeKey
	duplicateNodeRole
	duplicateInputKey
	failedToAddNodeFluxExecutors
	failedToAddInputFluxExecutors
	failedToEvaluateNodeFluxes
	failedToEvaluateInputFluxes
	nodeIsNotSubmittable
	requestDataIsNotAMap
	failedToAddSubmitFluxExecutors
	failedToRunGuards
	failedToSubmit

	invalidPropValue
)

var (
	// gex

	ErrFailedToGetExecutorHash  Grr = fmt.Errorf("grr%d", failedToGetExecutorHash)
	ErrInvalidValueType         Grr = fmt.Errorf("grr%d", invalidValueType)
	ErrKeyNotFound              Grr = fmt.Errorf("grr%d", keyNotFound)
	ErrIndexOutOfRange          Grr = fmt.Errorf("grr%d", indexOutOfRange)
	ErrFailedToEvaluateFluxData Grr = fmt.Errorf("grr%d", failedToEvaluateFluxData)

	// guards

	ErrInvalidLuaScript Grr = fmt.Errorf("grr%d", invalidLuaScript)

	// grove

	ErrFailedToValidateAndPopulateGrove Grr = fmt.Errorf("grr%d", failedToValidateAndPopulateGrove)
	ErrFailedToFindNode                 Grr = fmt.Errorf("grr%d", failedToFindNode)
	ErrFailedToRenderChild              Grr = fmt.Errorf("grr%d", failedToRenderChild)
	ErrInvalidNodeKey                   Grr = fmt.Errorf("grr%d", invalidNodeKey)
	ErrInvalidNodeType                  Grr = fmt.Errorf("grr%d", invalidNodeType)
	ErrInvalidInputKey                  Grr = fmt.Errorf("grr%d", invalidInputKey)
	ErrFailedToGetErrorNodeMessage      Grr = fmt.Errorf("grr%d", failedToGetErrorNodeMessage)
	ErrReceivedErrorNode                Grr = fmt.Errorf("grr%d", receivedErrorNode)
	ErrDuplicateNodeKey                 Grr = fmt.Errorf("grr%d", duplicateNodeKey)
	ErrDuplicateNodeRole                Grr = fmt.Errorf("grr%d", duplicateNodeRole)
	ErrDuplicateInputKey                Grr = fmt.Errorf("grr%d", duplicateInputKey)
	ErrFailedToAddNodeFluxExecutors     Grr = fmt.Errorf("grr%d", failedToAddNodeFluxExecutors)
	ErrFailedToAddInputFluxExecutors    Grr = fmt.Errorf("grr%d", failedToAddInputFluxExecutors)
	ErrFailedToEvaluateNodeFluxes       Grr = fmt.Errorf("grr%d", failedToEvaluateNodeFluxes)
	ErrFailedToEvaluateInputFluxes      Grr = fmt.Errorf("grr%d", failedToEvaluateInputFluxes)
	ErrNodeIsNotSubmittable             Grr = fmt.Errorf("grr%d", nodeIsNotSubmittable)
	ErrRequestDataIsNotAMap             Grr = fmt.Errorf("grr%d", requestDataIsNotAMap)
	ErrFailedToAddSubmitFluxExecutors   Grr = fmt.Errorf("grr%d", failedToAddSubmitFluxExecutors)
	ErrFailedToRunGuards                Grr = fmt.Errorf("grr%d", failedToRunGuards)
	ErrFailedToSubmit                   Grr = fmt.Errorf("grr%d", failedToSubmit)

	// grovex

	ErrInvalidEnumValue Grr = fmt.Errorf("grr%d", invalidPropValue)
)

func Wrap(grr Grr, format string, a ...any) error {
	return fmt.Errorf("[%w] "+format, append([]any{grr}, a...)...)
}
