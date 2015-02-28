package tasks

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	addr string

	tls  bool
	cert string
	key  string

	logger *log.Logger
}

type ServerOption func(*Server) error

func NewServer(opts ...ServerOption) (*Server, error) {
	s := &Server{
		addr: ":8000",
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func SetAddr(addr string) ServerOption {
	return func(s *Server) error {
		s.addr = addr
		return nil
	}
}

func SetTLS(cert, key string) ServerOption {
	return func(s *Server) error {
		s.cert = cert
		s.key = key
		s.tls = s.cert != "" && s.key != ""
		return nil
	}
}

func SetLogger(logger *log.Logger) ServerOption {
	return func(s *Server) error {
		s.logger = logger
		return nil
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	gs := grpc.NewServer()
	RegisterServiceServer(gs, s)
	if s.tls {
		s.printf("Listen on %s via TLS", s.addr)
		cred, err := credentials.NewServerTLSFromFile(s.cert, s.key)
		if err != nil {
			return err
		}
		return gs.Serve(cred.NewListener(lis))
	} else {
		s.printf("Listen on %s", s.addr)
		return gs.Serve(lis)
	}
}

func (s *Server) printf(format string, args ...interface{}) {
	if s.logger != nil {
		s.logger.Printf(format, args...)
	}
}

func (s *Server) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	res := &ListResponse{}
	res.Total = 1
	res.Tasks = make([]*Task, 1)
	res.Tasks[0] = &Task{Id: 1, Name: "Bring milk"}
	return res, nil
}
