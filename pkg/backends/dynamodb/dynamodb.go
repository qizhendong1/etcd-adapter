package dynamodb

import (
	"context"
	"github.com/k3s-io/kine/pkg/server"
)

// Options 包含 FoundationDB 的配置信息
type Options struct {
	Ip      string
	Port    int
	Table   string
	Timeout int
}

// fdbCache 是使用 FoundationDB 实现的后端缓存
//type fdbCache struct {
//	db  fdb.Database
//	dir directory.DirectorySubspace
//}

// NewFDBCache 返回一个实现了 server.Backend 接口的 FoundationDB 后端实例
func NewDynamoDBCache(ctx context.Context, options *Options) (server.Backend, error) {
	return nil, nil
}
