package srv

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type httpServer struct {
	server http.Server
}

func (h *httpServer) Start() error {
	err := h.server.ListenAndServe()
	log.Println("http server stop")
	return err
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
