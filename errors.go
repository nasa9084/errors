package errors

// WithError annotates error with another error. If e1 or e2 is nil, WithError returns nil.
func WithError(e1, e2 error) error {
	switch {
	case e1 == nil && e2 == nil:
		return nil
	case e1 == nil && e2 != nil:
		return e2
	case e1 != nil && e2 == nil:
		return e1
	}
	err := &withError{
		base:    e1,
		another: e2,
	}
	return err
}

type withError struct {
	base    error
	another error
}

func (w *withError) Error() string {
	return w.base.Error() + "/" + w.another.Error()
}

func (w *withError) Cause() error {
	return w.base
}
