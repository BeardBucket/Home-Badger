package hasser

import (
	"github.com/BeardBucket/Home-Badger/src/hasser/hzpub"
	"github.com/BeardBucket/Home-Badger/src/mainz/mzpub"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/pawal/go-hass"
)

type HassImpl struct {
	w   mzpub.Main
	evt hzpub.EventHass
}

type EventHassImpl struct {
	*HassImpl
	hAccess *hass.Access
	hCache  *cache.Cache[string]
}

func (h *HassImpl) NewEventHass() (hzpub.EventHass, error) {
	if h.evt != nil {
		return nil, hzpub.ErrEventHassExists
	}
	e := EventHassImpl{
		HassImpl: h,
	}
	h.evt = &e
	e.evt = &e
	if err := e.CreateAccess(); err != nil {
		h.evt = nil // We don't exist anymore
		return nil, err
	}
	return &e, nil
}

func (h *HassImpl) Evt() hzpub.EventHass {
	if h.evt == nil {
		if _, err := h.NewEventHass(); err != nil {
			h.Main().L().Error("Failed to create NewEventHass")
		}
	}
	return h.evt
}

func (h *HassImpl) Main() mzpub.Main {
	return h.w
}

func NewHass(w mzpub.Main) (hzpub.Hass, error) {
	h := HassImpl{
		w: w,
	}
	// Create all subtypes
	if _, err := h.NewEventHass(); err != nil {
		return nil, err
	}
	return &h, nil
}

// TestingF runs a quick, dev check - not for prod
func (h *HassImpl) TestingF() error { // TODO: Remove this
	a, err := h.Evt().Access()
	if err != nil {
		h.w.L().Warn("Failed to get API")
		return err
	}

	if err := a.CheckAPI(); err != nil {
		h.w.L().Warn("Failed API check")
		return err
	}
	h.w.L().Info("API ok")

	// Get the state of a device
	s, err := a.GetState("switch.relay_board_000_relay_01")
	if err != nil {
		return err
	}
	h.w.L().Info("Group %s state: %s\n", s.EntityID, s.State)

	// Create and interact with a device object
	dev, _ := a.GetDevice(s)
	h.w.L().Info("DEV: " + dev.EntityID())
	err = dev.On()
	if err != nil {
		return err
	}

	return nil
}
