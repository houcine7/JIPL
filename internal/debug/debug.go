package debug

type Error struct {
	Msg string
}

func (err *Error) Error() string {
	return err.Msg
}

func NewError(msg string) *Error {
	return &Error{Msg: msg}
}

var (
	NOERROR = &Error{Msg: ""}
)
