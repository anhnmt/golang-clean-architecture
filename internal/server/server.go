package server

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"connectrpc.com/vanguard/vanguardgrpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/rs/zerolog"
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

	srv *http.Server
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

	// Serve the http server on the http listener.
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", s.grpc.Host, s.grpc.Port)
		log.Info().Msgf("Starting application http://%s", addr)

		grpcServer := NewGrpcServer(s.grpc.LogPayload)
		transcoder, err := vanguardgrpc.NewTranscoder(grpcServer)
		if err != nil {
			return err
		}

		// create new http server
		s.srv = &http.Server{
			Addr: addr,
			// We use the h2c package in order to support HTTP/2 without TLS,
			// so we can handle gRPC requests, which requires HTTP/2, in
			// addition to Connect and gRPC-Web (which work with HTTP 1.1).
			Handler: h2c.NewHandler(
				transcoder,
				&http2.Server{},
			),
		}

		// run the server
		return s.srv.ListenAndServe()
	})

	return g.Wait()
}

func (s *Server) Close() error {
	return s.srv.Close()
}

// InterceptorLogger adapts zerolog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		log := l.With().Fields(fields).Logger()

		switch lvl {
		case logging.LevelDebug:
			log.Debug().Msg(msg)
		case logging.LevelInfo:
			log.Info().Msg(msg)
		case logging.LevelWarn:
			log.Warn().Msg(msg)
		case logging.LevelError:
			log.Error().Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func NewGrpcServer(logPayload bool) *grpc.Server {
	logger := InterceptorLogger(log.Logger)

	logEvents := []logging.LoggableEvent{
		logging.StartCall,
		logging.FinishCall,
	}

	// log payload if enabled
	if logPayload {
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

	return grpcServer
}
