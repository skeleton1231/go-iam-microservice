package storage

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/marmotedu/errors"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/log"
)

func singletonV2(cache bool) redis.UniversalClient {
	if cache {
		v := singleCachePool.Load()
		if v != nil {
			client, ok := v.(redis.UniversalClient)
			if !ok {
				log.Error("Stored value in singleCachePool is not of type redis.UniversalClient")
				return nil
			}
			return client
		}

		return nil
	}

	v := singlePool.Load()
	if v != nil {
		client, ok := v.(redis.UniversalClient)
		if !ok {
			log.Error("Stored value in singlePool is not of type redis.UniversalClient")
			return nil
		}
		return client
	}

	return nil
}

type RedisClusterV2 struct {
	KeyPrefix string
	HashKeys  bool
	IsCache   bool
	Ctx       context.Context
}

// ConnectToRedisV2 starts a go routine that periodically tries to connect to redis.
func ConnectToRedisV2(ctx context.Context, config *Config) {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	c := []RedisClusterV2{
		{},
		{IsCache: true, Ctx: context.Background()},
	}
	var ok bool
	for _, v := range c {
		if !connectSingletonV2(ctx, v.IsCache, config) {
			break
		}

		if !clusterConnectionIsOpenV2(ctx, v) {
			redisUp.Store(false)

			break
		}
		ok = true
	}
	redisUp.Store(ok)

again:
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			if !shouldConnect() {
				continue
			}
			for _, v := range c {
				if !connectSingletonV2(ctx, v.IsCache, config) {
					redisUp.Store(false)

					goto again
				}

				if !clusterConnectionIsOpenV2(ctx, v) {
					redisUp.Store(false)

					goto again
				}
			}
			redisUp.Store(true)
		}
	}
}

func connectSingletonV2(ctx context.Context, isCache bool, config *Config) bool {
	// NOTE: Assuming the connect logic remains the same, but if it differs, adjust accordingly.
	if singletonV2(isCache) == nil {
		log.Debug("Connecting to redis cluster")
		if isCache {
			singleCachePool.Store(NewRedisClusterPoolV2(isCache, config))
			return true
		}
		singlePool.Store(NewRedisClusterPoolV2(isCache, config))
		return true
	}
	return true
}

func clusterConnectionIsOpenV2(ctx context.Context, cluster RedisClusterV2) bool {
	c := singletonV2(cluster.IsCache)
	if c == nil {
		log.Warn("Redis client is nil")
		return false
	}

	testKey := "redis-test-" + uuid.Must(uuid.NewV4()).String()
	if err := c.Set(ctx, testKey, "test", time.Second).Err(); err != nil {
		log.Warnf("Error trying to set test key: %s", err.Error())
		return false
	}
	if _, err := c.Get(ctx, testKey).Result(); err != nil {
		log.Warnf("Error trying to get test key: %s", err.Error())
		return false
	}

	return true
}

// RedisOpts is the overridden type of redis.UniversalOptions. simple() and cluster() functions are not public in redis
// library.
// Therefore, they are redefined here to use in the creation of a new redis cluster logic.
// We don't want to use redis.NewUniversalClient() logic.
type RedisOptsV2 redis.UniversalOptions

func (opts *RedisOptsV2) simple() *redis.Options {
	return &redis.Options{
		Addr:         opts.Addrs[0],
		Password:     opts.Password,
		DB:           opts.DB,
		DialTimeout:  opts.DialTimeout,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		PoolSize:     opts.PoolSize,
		TLSConfig:    opts.TLSConfig,
	}
}

func (opts *RedisOptsV2) cluster() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:        opts.Addrs,
		Password:     opts.Password,
		DialTimeout:  opts.DialTimeout,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		PoolSize:     opts.PoolSize,
		TLSConfig:    opts.TLSConfig,
	}
}

func (opts *RedisOptsV2) failover() *redis.FailoverOptions {
	return &redis.FailoverOptions{
		MasterName:   opts.MasterName,
		Password:     opts.Password,
		DB:           opts.DB,
		DialTimeout:  opts.DialTimeout,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		PoolSize:     opts.PoolSize,
		TLSConfig:    opts.TLSConfig,
	}
}

// NewRedisClusterPool create a redis cluster pool.
func NewRedisClusterPoolV2(isCache bool, config *Config) redis.UniversalClient {
	// redisSingletonMu is locked and we know the singleton is nil
	log.Debug("Creating new Redis connection pool")

	// poolSize applies per cluster node and not for the whole cluster.
	poolSize := 500
	if config.MaxActive > 0 {
		poolSize = config.MaxActive
	}

	timeout := 5 * time.Second

	if config.Timeout > 0 {
		timeout = time.Duration(config.Timeout) * time.Second
	}

	var tlsConfig *tls.Config

	if config.UseSSL {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: config.SSLInsecureSkipVerify,
		}
	}

	var client redis.UniversalClient
	opts := &RedisOptsV2{
		Addrs:        getRedisAddrs(config),
		MasterName:   config.MasterName,
		Password:     config.Password,
		DB:           config.Database,
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		//IdleTimeout:  240 * timeout,
		PoolSize:  poolSize,
		TLSConfig: tlsConfig,
	}

	if opts.MasterName != "" {
		log.Info("--> [REDIS] Creating sentinel-backed failover client")
		client = redis.NewFailoverClient(opts.failover())
	} else if config.EnableCluster {
		log.Info("--> [REDIS] Creating cluster client")
		client = redis.NewClusterClient(opts.cluster())
	} else {
		log.Info("--> [REDIS] Creating single-node client")
		client = redis.NewClient(opts.simple())
	}

	return client
}

// Connect will establish a connection this is always true because we are dynamically using redis.
func (r *RedisClusterV2) Connect() bool {
	return true
}

func (r *RedisClusterV2) singleton() redis.UniversalClient {
	return singletonV2(r.IsCache)
}

func (r *RedisClusterV2) hashKey(in string) string {
	if !r.HashKeys {
		// Not hashing? Return the raw key
		return in
	}

	return HashStr(in)
}

func (r *RedisClusterV2) fixKey(keyName string) string {
	return r.KeyPrefix + r.hashKey(keyName)
}

func (r *RedisClusterV2) cleanKey(keyName string) string {
	return strings.Replace(keyName, r.KeyPrefix, "", 1)
}

func (r *RedisClusterV2) up() error {
	if !Connected() {
		return ErrRedisIsDown
	}

	return nil
}

// GetKey will retrieve a key from the database.
func (r *RedisClusterV2) GetKey(keyName string) (string, error) {
	if err := r.up(); err != nil {
		return "", err
	}

	cluster := r.singleton()

	value, err := cluster.Get(r.Ctx, r.fixKey(keyName)).Result()
	if err != nil {
		log.Debugf("Error trying to get value: %s", err.Error())

		return "", ErrKeyNotFound
	}

	return value, nil
}

func (r *RedisClusterV2) GetMultiKey(keys []string) ([]string, error) {
	if err := r.up(); err != nil {
		return nil, err
	}
	cluster := r.singleton()
	keyNames := make([]string, len(keys))
	copy(keyNames, keys)
	for index, val := range keyNames {
		keyNames[index] = r.fixKey(val)
	}

	result := make([]string, 0)

	switch v := cluster.(type) {
	case *redis.ClusterClient:
		getCmds := make([]*redis.StringCmd, len(keyNames))
		pipe := v.Pipeline()
		for i, key := range keyNames {
			getCmds[i] = pipe.Get(r.Ctx, key)
		}
		_, err := pipe.Exec(r.Ctx)
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Errorf("Error trying to get value: %s", err.Error())
			return nil, err
		}
		for _, cmd := range getCmds {
			result = append(result, cmd.Val())
		}

	case *redis.Client:
		values, err := v.MGet(r.Ctx, keyNames...).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Errorf("Error trying to get value: %s", err.Error())
			return nil, err
		}
		for _, val := range values {
			if val == nil {
				result = append(result, "")
			} else {
				result = append(result, fmt.Sprint(val))
			}
		}
	}

	notEmpty := false
	for _, val := range result {
		if val != "" {
			notEmpty = true
			break
		}
	}

	if notEmpty {
		return result, nil
	}
	return nil, ErrKeyNotFound
}
