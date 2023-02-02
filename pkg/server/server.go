package server

import "context"

// 抽象一个服务器接口，用于统一启停服务
type Server interface {
	// 启动服务
	Serve() error

	// 关闭服务
	Close(ctx context.Context) error
}
