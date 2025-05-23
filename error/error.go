package apperr

import (
	"context"
	crerr "github.com/cockroachdb/errors"
)

type AppErr struct {
	error
	Code       int    `json:"code"`
	Cause      string `json:"cause,omitempty"`
	Title      string `json:"title,omitempty"`
	Message    string `json:"message,omitempty"`
	wrappedErr error  `json:"-"`
}

func (a AppErr) Error() string {
	return a.wrappedErr.Error()
}

func New(cause string) *AppErr {
	wErr := crerr.New(cause)
	return &AppErr{
		Code:       500,
		Cause:      cause,
		wrappedErr: wErr,
	}
}

func Wrap(err error) *AppErr {
	e := crerr.Wrap(err, err.Error())
	return &AppErr{
		Code:       500,
		Cause:      err.Error(),
		wrappedErr: e,
	}
}

func (a *AppErr) WithCtx(ctx context.Context) *AppErr {
	a.wrappedErr = crerr.WithContextTags(a.wrappedErr, ctx)
	return a
}

func (a *AppErr) WithMessage(text string) *AppErr {
	a.wrappedErr = crerr.WithHint(a.wrappedErr, text)
	a.Message = text
	return a
}

func (a *AppErr) WithHttpStatus(httpStatus int) *AppErr {
	a.Code = httpStatus
	return a
}

func (a *AppErr) WithTitle(title string) *AppErr {
	a.Title = title
	return a
}

func (a *AppErr) WrappedError() error {
	return a.wrappedErr
}

type ErrorDef struct {
	Status  int
	Title   string
	Code    int
	Message string
}

func NewErrorDef(status int, title string, code int, message string) ErrorDef {
	return ErrorDef{
		Status:  status,
		Title:   title,
		Code:    code,
		Message: message,
	}
}

func (e ErrorDef) ToAppErr(err error) *AppErr {
	return Wrap(err).WithTitle(e.Title).WithMessage(e.Message).WithHttpStatus(e.Status)
}
