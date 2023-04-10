package hzpub

import "github.com/BeardBucket/Home-Badger/src/mainz/mzpub"

type EventHass interface {
	Hass
}

type Hass interface {
	Evt() EventHass
	Main() mzpub.Main
	NewEventHass() (EventHass, error)
	TestingF() error
}
