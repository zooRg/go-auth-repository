package todo

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	httpServe *http.Server
}

func (s *Server) Run(port string, handler http.Handler, logger *zap.SugaredLogger) error {
	logger.Debugf("server port: %s", port)

	s.httpServe = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServe.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServe.Shutdown(ctx)
}
