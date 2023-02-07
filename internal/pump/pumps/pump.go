package pumps

import (
	"context"
	"errors"
)

type Pump interface {
	GetName() string
	New() Pump
	Init(interface{}) error
	WriteData(context.Context, []interface{}) error
	SetFilters()
	GetFilters()
	SetTimeout(timeout int)
	GetTimeout() int
	SetOmitDetailedRecording(bool)
	GetOmitDetailedRecording() bool
}

// GetPumpByName returns the pump instance by given name.
func GetPumpByName(name string) (Pump, error) {
	if pump, ok := availablePumps[name]; ok && pump != nil {
		return pump, nil
	}

	return nil, errors.New(name + " Not found")
}
