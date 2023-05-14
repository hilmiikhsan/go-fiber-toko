package httpbin

import "context"

type HttpBinServiceInterface interface {
	PostMethod(ctx context.Context)
}
