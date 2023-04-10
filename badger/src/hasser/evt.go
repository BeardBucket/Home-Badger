package hasser

import "github.com/pawal/go-hass"

// CreateAccess creates and associated a hass.Access struct
func (i EventHassImpl) CreateAccess() error {
	a := hass.NewAccess("http://localhost:8123", "")
	err := a.CheckAPI()
	if err != nil {
		return err
	}
	i.w.L().Debug("API ok")
	i.hAccess = a
	return nil
}

// Access fetches the hass.Access struct
func (i EventHassImpl) Access() (*hass.Access, error) {
	if i.hAccess == nil {
		if err := i.CreateAccess(); err != nil {
			return nil, err
		}
	}
	return i.hAccess, nil
}
