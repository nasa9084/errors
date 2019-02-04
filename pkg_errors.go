package errors

import "github.com/pkg/errors"

// function aliases of github.com/pkg/errors
var (
	Cause        = errors.Cause
	Errorf       = errors.Errorf
	New          = errors.New
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	WithStack    = errors.WithStack
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
)

// type aliases of github.com/pkg/errors
type (
	Frame      = errors.Frame
	StackTrace = errors.StackTrace
)
