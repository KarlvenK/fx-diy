package server

import "golang.org/x/sync/errgroup"

type Server struct {
	Group errgroup.Group
	HttpS *HTTPServer
}

func NewServer() *Server {
	return &Server{}
}
