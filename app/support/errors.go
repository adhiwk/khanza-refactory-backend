package support

// ErrorKind jenis error bisnis; controller memetakannya ke HTTP status.
type ErrorKind int

const (
	KindNotFound ErrorKind = iota + 1
	KindConflict
	KindInvalid
	KindForbidden
)

// AppError error bisnis yang dikembalikan Action dan aman ditampilkan ke client.
type AppError struct {
	Kind    ErrorKind
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(message string) error  { return &AppError{Kind: KindNotFound, Message: message} }
func Conflict(message string) error  { return &AppError{Kind: KindConflict, Message: message} }
func Invalid(message string) error   { return &AppError{Kind: KindInvalid, Message: message} }
func Forbidden(message string) error { return &AppError{Kind: KindForbidden, Message: message} }
