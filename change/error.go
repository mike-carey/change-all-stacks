package change

type ErrorStack struct {
	message string
	errors []error
}

func (s *ErrorStack) Error() string {
	errStr := s.message + "\n-- Stack --\n"
	for _, e := range s.errors {
		errStr += " - " + e.Error() + "\n"
	}
	return errStr
}

func NewErrorStack(message string, errs []error) *ErrorStack {
	if len(errs) < 1 {
		return nil
	}

	return &ErrorStack{
		message: message,
		errors: errs,
	}
}
