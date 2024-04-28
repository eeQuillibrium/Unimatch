package grpcapp

import (
	"context"
	"fmt"
	"net"

	grpcserver "github.com/eeQuillibrium/Unimatch/auth_service/internal/app/grpc/server"
	"github.com/eeQuillibrium/Unimatch/auth_service/internal/service"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	"google.golang.org/grpc"
)

type App struct {
	log        *logger.Logger
	server     *grpc.Server
	serverPort int
}

func NewApp(
	log *logger.Logger,
	authService service.Auth,
	serverPort int,
) *App {
	server := grpc.NewServer()
	grpcserver.Register(server, authService)
	return &App{
		log:        log,
		server:     server,
		serverPort: serverPort,
	}
}

func (a *App) Run(ctx context.Context) error {
	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", a.serverPort))
	if err != nil {
		return err
	}
	a.log.Infof("server started on localhost:%d", a.serverPort)
	go func() {
		a.log.Fatal(a.server.Serve(lst))
	}()
	
	return nil
}
