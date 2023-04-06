package hasser

import (
	"fmt"
	"github.com/pawal/go-hass"
)

type EventHass interface {
	Test() error
}

type EventHassImpl struct {
}

func NewEventHass() (EventHass, error) {
	e := EventHassImpl{}

	return &e, nil
}

func (e EventHassImpl) Test() error {
	h := hass.NewAccess("http://localhost:8123", "")
	err := h.CheckAPI()
	if err != nil {
		panic(err)
	}
	fmt.Println("API ok")

	// Get the state of a device
	s, err := h.GetState("group.kitchen")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Group kitchen state: %s\n", s.State)

	// Create and interact with a device object
	dev, _ := h.GetDevice(s)
	fmt.Println("DEV: " + dev.EntityID())
	dev.On()

	return nil
}
