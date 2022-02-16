package cache

import "context"

type UserCache interface {
	GetUserCache(ctx context.Context,uid int64) (interface{},error)
}
