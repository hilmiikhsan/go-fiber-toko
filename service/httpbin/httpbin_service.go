package httpbin

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/client"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/model"
)

func NewHttpBinServiceInterface(httpBinClient *client.HttpBinClient) HttpBinServiceInterface {
	return &httpBinService{
		HttpBinClient: *httpBinClient,
	}
}

type httpBinService struct {
	client.HttpBinClient
}

func (h *httpBinService) PostMethod(ctx context.Context) {
	httpBin := model.HttpBin{
		Name: "ikhsan",
	}
	var response map[string]interface{}
	h.HttpBinClient.PostMethod(ctx, &httpBin, &response)
	common.NewLogger().Info("log response service ", response)
}
