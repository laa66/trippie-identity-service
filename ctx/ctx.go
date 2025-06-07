package ctx

import "context"

type Ctx interface {
	Context() context.Context
}