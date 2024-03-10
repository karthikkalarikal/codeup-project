package customerrors

type ErrorCode int

const (
	NoEmptyValueError ErrorCode = iota
	ValidatorError
	NoMatchingPasswordError
	JwtTokenMissingError
)

//go:generate stringer -type=ErrorCode
// func (e ErrorCode) String() string {
// 	switch e {
// 	case NoEmptyValueError:
// 		return "the fields cannot be left empty"
// 	case ValidatorError:
// 		return "the fields are not valid"
// 	case NoMatchingPasswordError:
// 		return "the passwords does not match"
// 	case JwtTokenMissingError:
// 		return "JWT token missing or invalid"
// 	default:
// 		return fmt.Sprintf("Unknown error code (%d)", e)
// 	}
// }
