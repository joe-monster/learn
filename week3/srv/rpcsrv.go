package srv

import (
	"google.golang.org/grpc"
	"net"
)

type rpcServer struct {
	server *grpc.Server
	port   string
}

func (h *rpcServer) Start() error {
	//time.Sleep(3 * time.Second)
	//return errors.New("rpc error")
	listen, err := net.Listen("tcp", ":"+h.port)
	if err != nil {
		return err
	}
	return h.server.Serve(listen)
}
func (h *rpcServer) Stop() error {
	h.server.Stop()
	return nil
}
func NewRpcServer(port string) *rpcServer {

	return &rpcServer{
		server: grpc.NewServer(),
		port:   port,
	}
}
