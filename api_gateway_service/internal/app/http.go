package app

import (
	"fmt"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
)

func (a *app) runHttpServer() error {
	a.echo.Server.ReadTimeout = readTimeout
	a.echo.Server.WriteTimeout = writeTimeout
	a.echo.Server.MaxHeaderBytes = maxHeaderBytes
	return a.echo.Start(fmt.Sprintf(":%d", a.cfg.Http.Port))
}
