package common

type genericError struct {
	err string
	d   string
}

func (e *genericError) Error() string {
	return e.err
}

func NewValidationError(text string, details string) error {
	return &genericError{"Input validation failed: " + text, details}
}

func NewDatabaseError(text string, details string) error {
	return &genericError{"DB-query failed: " + text, details}
}
