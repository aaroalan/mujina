package help

// PanicIfError : Helper method to avoid code duplication and simplify
// functions that need to validate an error.
func PanicIfError(error *error) {
	if *error != nil {
		panic(*error)
	}
}

// PanicIfFalse : Takes a boolean and message and panic the message if value is False.
func PanicIfFalse(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

// HasError : Returns True if the error pointer is not nil
func HasError(error *error) bool {
	return *error != nil
}
