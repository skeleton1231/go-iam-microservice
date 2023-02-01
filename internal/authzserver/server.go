package authzserver

import (
	"context"

	"github.com/skeleton1231/go-gin-restful-api-boilerplate/pkg/shutdown"

	genericoptions "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/options"
	genericapiserver "github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/server"
)

// RedisKeyPrefix defines the prefix key in redis for analytics data.
const RedisKeyPrefix = "analytics-"

type authzServer struct {
	gs               *shutdown.GracefulShutdown
	rpcServer        string
	clientCA         string
	redisOptions     *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
	analyticsOptions *analytics.AnalyticsOptions
	redisCancelFunc  context.CancelFunc
}

type preparedAuthzServer struct {
	*authzServer
}

// func createAuthzServer(cfg *config.Config) (*authzServer, error) {.
