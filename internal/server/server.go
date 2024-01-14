package server

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"connectrpc.com/vanguard/vanguardgrpc"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	gen "github.com/anhnmt/golang-clean-architecture/api/protobuf/gen"
	"github.com/anhnmt/golang-clean-architecture/pkg/config"
	"github.com/anhnmt/golang-clean-architecture/pkg/postgres"
)

var _ Server = (*server)(nil)

type server struct {
	cfg config.Server

	pg         postgres.DBEngine
	grpcServer *grpc.Server
}

func New(
	cfg config.Server,
	pg postgres.DBEngine,
	grpcServer *grpc.Server,

	_ gen.UserServiceServer,
) Server {
	return &server{
		cfg:        cfg,
		pg:         pg,
		grpcServer: grpcServer,
	}
}

func (s *server) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	if *s.cfg.Pprof.Enable {
		g.Go(func() error {
			addr := fmt.Sprintf("%s:%d", s.cfg.Pprof.Host, s.cfg.Pprof.Port)
			log.Info().Msgf("Starting pprof http://%s", addr)

			return http.ListenAndServe(addr, nil)
		})
	}

	// Serve the http server on the http listener.
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", s.cfg.Grpc.Host, s.cfg.Grpc.Port)
		log.Info().Msgf("Starting application http://%s", addr)

		transcoder, err := vanguardgrpc.NewTranscoder(s.grpcServer)
		if err != nil {
			return err
		}

		// create new http server
		srv := &http.Server{
			Addr: addr,
			// We use the h2c package in order to support HTTP/2 without TLS,
			// so we can handle gRPC requests, which requires HTTP/2, in
			// addition to Connect and gRPC-Web (which work with HTTP 1.1).
			Handler: h2c.NewHandler(
				transcoder,
				&http2.Server{},
			),
		}

		srv.Close()

		// run the server
		return srv.ListenAndServe()
	})

	return g.Wait()
}

func (s *server) Close() {
	s.grpcServer.GracefulStop()
}
