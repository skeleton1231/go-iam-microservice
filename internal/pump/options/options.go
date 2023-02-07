package options

import (
	"github.com/marmotedu/log"

	genericoptions "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/options"
)

// PumpConfig defines options for pump back-end.
type PumpConfig struct {
	Type                  string                     `json:"type"                    mapstructure:"type"`
	Filters               analytics.AnalyticsFilters `json:"filters"                 mapstructure:"filters"`
	Timeout               int                        `json:"timeout"                 mapstructure:"timeout"`
	OmitDetailedRecording bool                       `json:"omit-detailed-recording" mapstructure:"omit-detailed-recording"`
	Meta                  map[string]interface{}     `json:"meta"                    mapstructure:"meta"`
}

// Options runs a pumpserver.
type Options struct {
	PurgeDelay            int                          `json:"purge-delay"             mapstructure:"purge-delay"`
	Pumps                 map[string]PumpConfig        `json:"pumps"                   mapstructure:"pumps"`
	HealthCheckPath       string                       `json:"health-check-path"       mapstructure:"health-check-path"`
	HealthCheckAddress    string                       `json:"health-check-address"    mapstructure:"health-check-address"`
	OmitDetailedRecording bool                         `json:"omit-detailed-recording" mapstructure:"omit-detailed-recording"`
	RedisOptions          *genericoptions.RedisOptions `json:"redis"                   mapstructure:"redis"`
	Log                   *log.Options                 `json:"log"                     mapstructure:"log"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	s := Options{
		PurgeDelay: 10,
		Pumps: map[string]PumpConfig{
			"csv": {
				Type: "csv",
				Meta: map[string]interface{}{
					"csv_dir": "./analytics-data",
				},
			},
		},
		HealthCheckPath:    "healthz",
		HealthCheckAddress: "0.0.0.0:7070",
		RedisOptions:       genericoptions.NewRedisOptions(),
		Log:                log.NewOptions(),
	}

	return &s
}
