package hasser

import (
	"github.com/pawal/go-hass"
	"os"
)

func init() {

}

// CreateAccess creates and associated a hass.Access struct
func (i *EventHassImpl) CreateAccess() error {
	a := hass.NewAccess("http://supervisor", "")
	a.SetBearerToken(os.Getenv("SUPERVISOR_TOKEN"))
	a.SetPath(hass.PathTypeAPI, "/")
	if err := a.CheckAPI(); err != nil {
		return err
	}
	i.w.L().Debug("API ok")
	i.hAccess = a
	return nil
}

// Access fetches the hass.Access struct
func (i *EventHassImpl) Access() (*hass.Access, error) {
	if i.hAccess == nil {
		if err := i.CreateAccess(); err != nil {
			return nil, err
		}
	}
	return i.hAccess, nil
}
