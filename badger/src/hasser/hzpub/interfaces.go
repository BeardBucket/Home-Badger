package hzpub

import (
	"github.com/BeardBucket/Home-Badger/src/mainz/mzpub"
	"github.com/pawal/go-hass"
)

type EventHass interface {
	Hass
	CreateAccess() error
	Access() (*hass.Access, error)
}

type Hass interface {
	Evt() EventHass
	Main() mzpub.Main
	NewEventHass() (EventHass, error)
	TestingF() error
}
