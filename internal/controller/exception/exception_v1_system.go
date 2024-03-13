package exception

import (
	"context"
	"myapp/internal/service"

	"myapp/api/exception/v1"
)

func (c *ControllerV1) System(ctx context.Context, req *v1.SystemReq) (res *v1.SystemRes, err error) {
	err = service.Exception().System()
	if err != nil {
		return nil, err
	}
	return &v1.SystemRes{
		Name: "system",
		Age:  1,
	}, nil
}
