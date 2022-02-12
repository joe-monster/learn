package srv

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpServer struct {
	server http.Server
}

func (h *httpServer) Start() error {
	return h.server.ListenAndServe()
}
func (h *httpServer) Stop() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return h.server.Shutdown(ctx)
}
func NewHttpServer(port string) *httpServer {
	router := gin.Default()
	router.Use(gin.Recovery())

	//插件和路由...

	return &httpServer{
		server: http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
	}
}
