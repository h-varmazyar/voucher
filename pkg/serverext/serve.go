package serverext

import (
	"github.com/h-varmazyar/voucher/pkg/netext"
	log "github.com/sirupsen/logrus"
	"net"
)

type ServeFunc func(listener net.Listener) error

func (serve ServeFunc) listen(port netext.Port) error {
	listener, err := net.Listen(netext.TCP, netext.LocalAddress(port))
	if err != nil {
		return err
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Errorf("failed to close network listener: %v", err)
		}
	}()

	return serve(listener)
}
