package server

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

type Server struct {
	pprof config.Pprof
	grpc  config.Grpc
}

func New(cfg config.Server) *Server {
	return &Server{
		pprof: cfg.Pprof,
		grpc:  cfg.Grpc,
	}
}

func (s *Server) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	if s.pprof.Enable {
		g.Go(func() error {
			addr := fmt.Sprintf("%s:%d", s.pprof.Host, s.pprof.Port)
			log.Info().Msgf("Starting pprof http://%s", addr)

			return http.ListenAndServe(addr, nil)
		})
	}

	logger := InterceptorLogger(log.Logger)

	logEvents := []logging.LoggableEvent{
		logging.StartCall,
		logging.FinishCall,
	}

	// log payload if enabled
	if s.grpc.LogPayload {
		logEvents = append(logEvents,
			logging.PayloadReceived,
			logging.PayloadSent,
		)
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logEvents...),
	}

	streamInterceptors := []grpc.StreamServerInterceptor{
		logging.StreamServerInterceptor(logger, opts...),
		recovery.StreamServerInterceptor(),
		validator.StreamServerInterceptor(),
	}
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		logging.UnaryServerInterceptor(logger, opts...),
		recovery.UnaryServerInterceptor(),
		validator.UnaryServerInterceptor(),
	}

	// register grpc service Server
	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(streamInterceptors...),
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)

	// Serve the http server on the http listener.
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", s.grpc.Host, s.grpc.Port)
		log.Info().Msgf("Starting application http://%s", addr)

		// create new http server
		srv := &http.Server{
			Addr: addr,
			// Use h2c, so we can serve HTTP/2 without TLS.
			Handler: h2c.NewHandler(
				grpcServer,
				&http2.Server{},
			),
			// ReadHeaderTimeout: 10 * time.Second,
			// ReadTimeout:       1 * time.Minute,
			// WriteTimeout:      1 * time.Minute,
			// MaxHeaderBytes:    8 * 1024, // 8KiB
		}

		// run the server
		return srv.ListenAndServe()
	})

	return g.Wait()
}
