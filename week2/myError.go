package week2

var notFoundCode = 40001
var sysErrorCode = 50001

type myError struct {
	code int
	msg  string
}

func (err myError) Error() string {
	return err.msg
}

func IsNotFoundError(err error) bool {
	e,ok := err.(myError)
	if ok {
		return e.code == notFoundCode
	} else {
		return false
	}
}
