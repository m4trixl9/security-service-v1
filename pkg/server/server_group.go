package server

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// 定义一个服务器Group，将所有的服务都放到该组里，如：各种http服务、rpc服务等
type ServerGroup struct {
	servers []Server
	logger  *zap.Logger
}

func NewServerGroup(svr ...Server) *ServerGroup {
	return &ServerGroup{
		servers: svr,
		logger:  nil,
	}
}

// Run
// 启动 ServerGroup 中的所有服务，调用 Run 将阻塞，直到 ctx 通过 Context.Done 触发服务停止
func (g *ServerGroup) Run(ctx context.Context) {
	// 用 sync.WaitGroup 来做并发控制
	var wg sync.WaitGroup

	wg.Add(len(g.servers))
	for _, srv := range g.servers {
		// 用 goroutine 来启动一个常驻服务
		go func(s Server) {
			defer wg.Done()

			if err := s.Serve(); err != nil {
				//
				fmt.Println("server serve error" + err.Error())
			}
		}(srv)
	}

	// 用 goroutine 来接收由ctx带过来的中断，然后做优雅关闭（做各服务的收尾工作），各个服务的优雅关闭由统一封装的Server各自实现
	go func() {
		// 阻塞
		// channel的规则：当写满时，不可以写；取空时，不可以取；发送数据将持续阻塞，直到数据被接收；接收将持续阻塞，直到发送方发送数据，每次只接收一个元素；
		// 读一个read only的channel，除非该channel被关闭，否则读不出来任何东西；如果被关闭，则会读出相应类型的零值
		<-ctx.Done()
		for _, srv := range g.servers {
			srv.Close(context.Background())
		}
	}()

	// 阻塞，直到计数器为0
	wg.Wait()
}
