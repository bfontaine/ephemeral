package ephemeral

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/netutil"

	"github.com/bfontaine/ephemeral/Godeps/_workspace/src/github.com/hydrogen18/stoppableListener"
)

type Server struct {
	http    http.Server
	sl      *stoppableListener.StoppableListener
	data    interface{}
	stopped bool
}

func New() *Server {
	return &Server{}
}

func (s *Server) Stop(data interface{}) {
	log.Println("stop")
	if s.stopped {
		return
	}
	log.Println("    ...")
	s.data = data
	s.sl.Stop()
	s.stopped = true
}

func (s *Server) HandleFunc(path string,
	fn func(*Server, http.ResponseWriter, *http.Request)) {

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fn(s, w, r)
	})
}

func (s *Server) Listen(host string) (data interface{}, err error) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return
	}

	s.sl, err = stoppableListener.New(listener)
	if err != nil {
		return
	}

	defer func() { data = s.data }()

	s.http.Serve(netutil.LimitListener(s.sl, 1))

	return
}
