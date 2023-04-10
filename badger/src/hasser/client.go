package hasser

import (
	"github.com/BeardBucket/Home-Badger/src/hasser/hzpub"
	"github.com/BeardBucket/Home-Badger/src/mainz/mzpub"
	"github.com/pawal/go-hass"
)

type HassImpl struct {
	w   mzpub.Main
	evt hzpub.EventHass
}

type EventHassImpl struct {
	HassImpl
}

func (h HassImpl) NewEventHass() (hzpub.EventHass, error) {
	e := EventHassImpl{
		HassImpl: h,
	}
	h.evt = &e
	return &e, nil
}

func (h HassImpl) Evt() hzpub.EventHass {
	return h.evt
}
func (h HassImpl) Main() mzpub.Main {
	return h.w
}

func NewHass(w mzpub.Main) (hzpub.Hass, error) {
	h := HassImpl{
		w: w,
	}
	if _, err := h.NewEventHass(); err != nil {
		return nil, err
	}
	return &h, nil
}

// TestingF runs a quick, dev check - not for prod
func (h HassImpl) TestingF() error { // TODO: Remove this
	a := hass.NewAccess("http://localhost:8123", "")
	err := a.CheckAPI()
	if err != nil {
		return err
	}
	h.w.L().Info("API ok")

	// Get the state of a device
	s, err := a.GetState("group.kitchen")
	if err != nil {
		return err
	}
	h.w.L().Info("Group kitchen state: %s\n", s.State)

	// Create and interact with a device object
	dev, _ := a.GetDevice(s)
	h.w.L().Info("DEV: " + dev.EntityID())
	err = dev.On()
	if err != nil {
		return err
	}

	return nil
}
