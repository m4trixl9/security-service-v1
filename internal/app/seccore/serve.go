package seccore

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/device-security-v1/seccore/pkg/server"
)

// Serve
// 运行seccore服务，内部实现优雅关闭，外部可通过 ctx 来控制生命周期
// 优雅关闭：在收到 interrupt 或者 terminate 信号时，触发服务的主动关闭
// 该函数最终在 secserve cobra命令中调用，即这里是secserve命令的handler
func Serve(ctx context.Context) {

	// 配置初始化
	initServiceWithConfig()

	// 将所有的服务都添加到ServerGroup中，以便于进行统一启停
	g := server.NewServerGroup(
		httpServer(),
	)

	// 监控系统中断信号、创建Context：ctx
	// 在收到信号时，会自动触发ctx的Done
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	g.Run(ctx)
}
