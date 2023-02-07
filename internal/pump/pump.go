package pump

import (
	"context"
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
