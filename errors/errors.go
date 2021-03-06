package errors

import "strconv"

// Error error.
type Error interface {
	error
	// Code get error code.
	Code() int
	Equal(error) bool
}

// ecode error.
type ecode int

var (
	OK           ecode
	NotModified  ecode = -304
	ParamsErr    ecode = -400
	Unauthorized ecode = -401
	NothingFound ecode = -404
	ServerErr    ecode = -500
)

func (e ecode) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e ecode) Code() int {
	return int(e)
}

func (e ecode) Equal(err error) bool {
	cd := Code(err)
	return e.Code() == cd.Code()
}

// Code converts error to ecode.
func Code(e error) (ie Error) {
	if e == nil {
		ie = OK
		return
	}
	i, err := strconv.Atoi(e.Error())
	if err != nil {
		i = -500
	}
	ie = ecode(i)
	return
}

// Int converts int to ecode.
func Int(i int) (ie Error) {
	return ecode(i)
}
