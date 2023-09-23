package errors

type Error struct {
	Message string
}

func New(s string) error {
	return &Error{Message: s}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrExtFormatNotSupported  = New("extension format not supported")
	ErrFailedOpenSrcFile      = New("failed to open source file")
	ErrNotHaveExt             = New("file doesn't have any extension")
	ErrImgDecoder             = New("error while decode image")
	ErrFailedCreateOutputDir  = New("failed to create dir")
	ErrFailedReadOutputDir    = New("failed to read output directory")
	ErrFailedCreateOutputFile = New("failed to create output file")
	ErrImgEncoder             = New("error while encode image")
)
