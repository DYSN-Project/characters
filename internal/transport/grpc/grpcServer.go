package grpc

import (
	"dysn/character/internal/manager"
	"dysn/character/internal/service/logger"
	"dysn/character/internal/service/validation"
	"dysn/character/internal/transport/grpc/pb"
	"google.golang.org/grpc"
	"net"
)
type GrpcServerInterface interface {
	StartServer()
	StopServer()
}

type GrpcServer struct {
	server *grpc.Server
	port   string
	logger *logger.Logger
}

func NewGrpcServer(port string,
	logger *logger.Logger,
	vld validation.ValidationInterface,
	mng manager.CharacterManagerInterface) *GrpcServer {
	srv := grpc.NewServer()
	chrSrv := NewCharacterServer(logger,vld,mng)
	pb.RegisterCharacterServer(srv, chrSrv)

	return &GrpcServer{server: srv, port: port, logger: logger}
}

func (g *GrpcServer) StartServer() {
	g.logger.InfoLog.Println("Character server starting...")
	l, err := net.Listen("tcp", g.port)
	if err != nil {
		g.logger.ErrorLog.Panic(err)
	}
	err = g.server.Serve(l)
	if err != nil {
		g.logger.ErrorLog.Panic(err)
	}
}

func (g *GrpcServer) StopServer() {
	g.logger.InfoLog.Println("Character server stopping...")
	g.server.Stop()
}