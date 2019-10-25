package cache

type stateNotFound struct {
	error string
}

func (err *stateNotFound) Error() string {
	return err.error
}
