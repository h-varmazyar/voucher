package serverext

import (
	"github.com/h-varmazyar/voucher/pkg/netext"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	serves map[netext.Port]ServeFunc
	logger *log.Logger
}

func New(logger *log.Logger) *Server {
	service := new(Server)

	service.serves = make(map[netext.Port]ServeFunc)
	service.logger = logger

	return service
}

func (s *Server) Serve(port netext.Port, function ServeFunc) {
	s.serves[port] = function
}

func (s *Server) Start(name, version string) {
	var exposes []netext.Port
	interrupt := make(chan error)
	for port, serveFunc := range s.serves {
		go func(port netext.Port, serveFunc ServeFunc) {
			interrupt <- serveFunc.listen(port)
		}(port, serveFunc)
	}

	fields := log.Fields{
		"service": name,
		"version": version,
		"exposes": exposes,
	}

	log.WithFields(fields).Info("running service...")

	interruptErr := <-interrupt

	//todo: close ports

	if interruptErr == nil {
		log.WithFields(fields).Panic("service interrupted")
	} else {
		log.WithFields(fields).WithError(interruptErr).Fatal("service stopped")
	}
}
