package srv

import (
	"google.golang.org/grpc"
	"log"
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
	err = h.server.Serve(listen)
	log.Println("rpc server stop")
	return err
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
