package app

import (
	"context"
	"time"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

const (
	maxHeaderBytes = 1 << 20
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
)

func (a *app) runHttpServer(ctx context.Context) error {
	a.echo.Server.ReadTimeout = readTimeout
	a.echo.Server.WriteTimeout = writeTimeout
	a.echo.Server.MaxHeaderBytes = maxHeaderBytes
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithDomain("up-distinctly-chamois.ngrok-free.app"),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}
	a.echo.Listener = listener
	return a.echo.Start("")
}
