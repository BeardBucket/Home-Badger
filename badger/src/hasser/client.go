package hasser

import (
	"github.com/BeardBucket/Home-Badger/src/mainz"
	"github.com/pawal/go-hass"
)

type EventHass interface {
	TestingF() error
}

type EventHassImpl struct {
	w mainz.Main
}

func NewEventHass(w mainz.Main) (EventHass, error) {
	e := EventHassImpl{
		w: w,
	}

	return &e, nil
}

// TestingF runs a quick, dev check - not for prod
func (e EventHassImpl) TestingF() error { // TODO: Remove this
	h := hass.NewAccess("http://localhost:8123", "")
	err := h.CheckAPI()
	if err != nil {
		return err
	}
	e.w.L().Info("API ok")

	// Get the state of a device
	s, err := h.GetState("group.kitchen")
	if err != nil {
		return err
	}
	e.w.L().Info("Group kitchen state: %s\n", s.State)

	// Create and interact with a device object
	dev, _ := h.GetDevice(s)
	e.w.L().Info("DEV: " + dev.EntityID())
	err = dev.On()
	if err != nil {
		return err
	}

	return nil
}
