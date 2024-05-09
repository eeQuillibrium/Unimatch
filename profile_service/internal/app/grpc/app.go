package grpcapp

import (
	"fmt"
	"net"

	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	grpcserver "github.com/eeQuillibrium/Unimatch/profile_service/internal/app/grpc/server"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/service"
	"google.golang.org/grpc"
)

type gRPCApp struct {
	log        *logger.Logger
	cfg        *config.Config
	grpcServer *grpc.Server
}

func NewGRPCApp(
	log *logger.Logger,
	cfg *config.Config,
	service service.ProfileProvider,
) *gRPCApp {
	grpcServer := grpc.NewServer()
	grpcserver.Register(grpcServer, service)

	return &gRPCApp{log: log, cfg: cfg, grpcServer: grpcServer}
}

func (a *gRPCApp) Run() error {
	a.log.Info("run grpc server...")

	lst, err := net.Listen("tcp", fmt.Sprintf("%d", a.cfg.GRPC.Serverport))
	if err != nil {
		return err
	}

	go func() {
		a.log.Fatal(a.grpcServer.Serve(lst))
	}()

	return nil
}
