package options

import (
	"github.com/marmotedu/log"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/authzserver/analytics"
	genericoptions "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/options"
)

// Options runs a authzserver.
type Options struct {
	RPCServer               string                                 `json:"rpcserver"      mapstructure:"rpcserver"`
	ClientCA                string                                 `json:"client-ca-file" mapstructure:"client-ca-file"`
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"         mapstructure:"server"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure"       mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"         mapstructure:"secure"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"          mapstructure:"redis"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"        mapstructure:"feature"`
	Log                     *log.Options                           `json:"log"            mapstructure:"log"`
	AnalyticsOptions        *analytics.AnalyticsOptions            `json:"analytics"      mapstructure:"analytics"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		RPCServer:               "127.0.0.1:8081",
		ClientCA:                "",
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		RedisOptions:            genericoptions.NewRedisOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
		Log:                     log.NewOptions(),
		AnalyticsOptions:        analytics.NewAnalyticsOptions(),
	}

	return &o
}
