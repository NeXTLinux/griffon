package griffon

import (
	"github.com/nextlinux/griffon/internal/bus"
	"github.com/nextlinux/griffon/internal/log"
	"github.com/wagoodman/go-partybus"

	"github.com/anchore/go-logger"
)

func SetLogger(logger logger.Logger) {
	log.Log = logger
}

func SetBus(b *partybus.Bus) {
	bus.SetPublisher(b)
}
