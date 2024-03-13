// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IException interface {
		Business() error
		System() error
	}
)

var (
	localException IException
)

func Exception() IException {
	if localException == nil {
		panic("implement not found for interface IException, forgot register?")
	}
	return localException
}

func RegisterException(i IException) {
	localException = i
}
