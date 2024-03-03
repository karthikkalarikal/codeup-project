package customerrors

const (
	NoEmptyValueError       = "the fields cannot be left empty"
	ValidatorError          = "the fields are not valid"
	NoMatchingPasswordError = "the passwords doesnot match"
	JwtTokenMissingError    = "JWT token missing or invalid"
)

// type ConstantErrors struct {
// }

// func (e *ConstantErrors) NoEmptyValueError() string {
// 	return "the fields cannot be left empty"
// }
