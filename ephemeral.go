package ephemeral

import (
	"net"
	"net/http"

	"github.com/bfontaine/ephemeral/Godeps/_workspace/src/github.com/hydrogen18/stoppableListener"
)

// Server is a wrapper for an ephemeral http.Server.
type Server struct {
	http    http.Server
	sl      *stoppableListener.StoppableListener
	data    interface{}
	stopped bool
}

// Handler represents the type of functions used as handlers with an ephemeral
// Server
type Handler func(*Server, http.ResponseWriter, *http.Request)

// New returns a pointer on a new Server
func New() *Server {
	return &Server{}
}

// Stop stops the server and use its argument as a return value for Listen.
func (s *Server) Stop(data interface{}) {
	if s.stopped {
		return
	}
	s.data = data
	s.sl.Stop()
	s.stopped = true
}

// HandleFunc adds a new handler for a given path
func (s *Server) HandleFunc(path string, fn Handler) {

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fn(s, w, r)
	})
}

// Listen works like http.ListenAndServe but returns the argument passed to
// Stop.
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

	s.http.Serve(s.sl)

	return
}

// GetRequest is a shortcut function which creates an ephemeral Server, run it,
// and stops once it got one request, which it then returns.
func GetRequest(host, path string) (*http.Request, error) {
	s := New()
	s.HandleFunc(path, func(s *Server, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		s.Stop(r)
	})
	r, err := s.Listen(host)
	return r.(*http.Request), err
}
