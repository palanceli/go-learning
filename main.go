package main

import (
	"github.com/golang/glog"
)

type Server struct {
	x, y int
}

func (s *Server) Start() {
	glog.Info("start server...")
}

type Proc struct {
	z int
}

func (p *Proc) Start() {
	glog.Info("start proc...")
}

type Instance struct {
	x, y   int
	server *Server
	proc   *Proc
}

type Service interface {
	Start()
}

func StartService(s Service) {
	s.Start()
}

func main() {
	defer glog.Flush()
	obj := &Instance{
		x: 0,
		y: 1,
		server: &Server{
			x: 2,
			y: 3,
		},
		proc: &Proc{
			z: 4,
		},
	}
	StartService(obj.server)
}
