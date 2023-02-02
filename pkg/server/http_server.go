package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	svr *http.Server
}

// 将http.Server封装成一个自定义的httpServer，方便ServerGroup处理
func WrapHTTP(svr *http.Server) Server {
	return &httpServer{
		svr: svr,
	}
}

// 进一步封装，将Addr & handler封装到Gin Engine中
func WrapGin(e *gin.Engine, Addr string) Server {
	return WrapHTTP(&http.Server{
		Addr:    Addr,
		Handler: e,
	})
}

func (h *httpServer) Serve() error {
	if err := h.svr.ListenAndServe(); err != nil {
		//
		if err == http.ErrServerClosed {
			return err ///////
		}
		return err
	}
	return nil
}

func (h *httpServer) Close(ctx context.Context) error {
	return h.svr.Shutdown(ctx)
}

// interface guard
var _ Server = (*httpServer)(nil)
