package exception

import (
	"context"
	"myapp/internal/service"

	"myapp/api/exception/v1"
)

func (c *ControllerV1) Business(ctx context.Context, req *v1.BusinessReq) (res *v1.BusinessRes, err error) {
	err = service.Exception().Business()
	if err != nil {
		return nil, err
	}
	return &v1.BusinessRes{
		Name: "business",
		Age:  1,
	}, nil
}
