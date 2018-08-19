package hashdb

import "errors"

var (
	// 创建数据库错误
	ErrCreatingDatabase = errors.New(`error while creating Database`)
	// 没有找到数据记录
	ErrDataStoreNotFound = errors.New(`the requested data store was not found`)
)
