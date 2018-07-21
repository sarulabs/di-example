package helpers

// ErrValidation is the error type that should be used
// to indicate that the error is caused by a validation problem.
type ErrValidation struct {
	msg string
}

// NewErrValidation is the ErrValidation constructor.
func NewErrValidation(msg string) *ErrValidation {
	return &ErrValidation{msg: msg}
}

// Error returns the error message.
func (err *ErrValidation) Error() string {
	return err.msg
}

// ErrNotFound is the error type that should be used
// to indicate that the error is caused by the nonexistence of the requested resource.
type ErrNotFound struct {
	msg string
}

// NewErrNotFound is the ErrValidation constructor.
func NewErrNotFound(msg string) *ErrNotFound {
	return &ErrNotFound{msg: msg}
}

// Error returns the error message.
func (err *ErrNotFound) Error() string {
	return err.msg
}
