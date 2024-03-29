package grpc_server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

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

func New(cfg config.Grpc) *grpc.Server {
	logger := InterceptorLogger(log.Logger)

	logEvents := []logging.LoggableEvent{
		logging.StartCall,
		logging.FinishCall,
	}

	// log payload if enabled
	if *cfg.LogPayload {
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

	// register grpc service server
	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(streamInterceptors...),
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)
	reflection.Register(grpcServer)

	return grpcServer
}
