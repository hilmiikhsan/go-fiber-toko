package client

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{})
}
