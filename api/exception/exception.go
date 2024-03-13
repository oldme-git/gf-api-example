// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package exception

import (
	"context"

	"myapp/api/exception/v1"
)

type IExceptionV1 interface {
	Business(ctx context.Context, req *v1.BusinessReq) (res *v1.BusinessRes, err error)
	System(ctx context.Context, req *v1.SystemReq) (res *v1.SystemRes, err error)
}
