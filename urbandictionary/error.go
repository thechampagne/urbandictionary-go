package urbandictionary

type Error struct{
	Message string
}

func (err Error) Error() string {
	return err.Message
}

func newError(message string) error {
	return Error{message}
}