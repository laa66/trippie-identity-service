package apperr

import (
	"context"
	crerr "github.com/cockroachdb/errors"
)

// TODO: Add err definition
type AppErr struct {
	error
	Code       int    `json:"code"`
	Cause      string `json:"cause,omitempty"`
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

func (a *AppErr) WrappedError() error {
	return a.wrappedErr
}
