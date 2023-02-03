package load

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/marmotedu/log"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/pkg/storage"
)

// Loader defines function to reload storage.
type Loader interface {
	Reload() error
}

// Load is used to reload given storage.
type Load struct {
	ctx    context.Context
	lock   *sync.RWMutex
	loader Loader
}

// NewLoader return a loader with a loader implement.
func NewLoader(ctx context.Context, loader Loader) *Load {
	return &Load{
		ctx:    ctx,
		lock:   new(sync.RWMutex),
		loader: loader,
	}
}

// Start start a loop service.
func (l *Load) Start() {
	go startPubSubLoop()
	go l.reloadQueueLoop()
	// 1s is the minimum amount of time between hot reloads. The
	// interval counts from the start of one reload to the next.
	go l.reloadLoop()
	l.DoReload()
}

func startPubSubLoop() {
	cacheStore := storage.RedisCluster{}
	cacheStore.Connect()
	// On message, synchronize
	for {
		err := cacheStore.StartPubSubHandler(RedisPubSubChannel, func(v interface{}) {
			handleRedisEvent(v, nil, nil)
		})
		if err != nil {
			if !errors.Is(err, storage.ErrRedisIsDown) {
				log.Errorf("Connection to Redis failed, reconnect in 10s: %s", err.Error())
			}

			time.Sleep(10 * time.Second)
			log.Warnf("Reconnecting: %s", err.Error())
		}
	}
}
