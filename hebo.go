package hebo

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/Presbyter/hebo/pkg/log"
)

type Server struct {
	cfg        *Config
	log        log.Logger
	httpServer *http.Server
}

func New(c *Config) *Server {
	return &Server{
		cfg:        c,
		log:        log.Default,
		httpServer: &http.Server{},
	}
}

func (s *Server) SetLogger(l log.Logger) {
	if l == nil {
		s.log = log.Default
		return
	}
	s.log = l
}

func (s *Server) Run(ctx context.Context) {
	go s.serve()

	go func() {
		<-ctx.Done()
		s.httpServer.Shutdown(ctx)
		s.log.Debugf("exit server now")
	}()
}

func (s *Server) serve() {
	server := s.httpServer
	server.Addr = fmt.Sprintf("%s:%d", s.cfg.IpBind, s.cfg.Port)
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		s.log.Errorf("create tcp listener fail. error: %s", err)
		return
	}

	upstreamService := InitUpStream(s.cfg)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := upstreamService.Forward(w, r); err != nil {
			s.log.Errorf("upstream forward request fail. error: %s", err)
		}
	})

	if err := server.Serve(listener); err != nil {
		s.log.Errorf("start http server fail. error: %s", err)
		return
	}
}
