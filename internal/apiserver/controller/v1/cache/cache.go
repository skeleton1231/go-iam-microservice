// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package cache defines a cache service which can return all secrets and policies.
package cache

import (
	"fmt"
	"sync"

	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/apiserver/store"
)

// Cache defines a cache service used to list all secrets and policies.
type Cache struct {
	store store.Factory
}

var (
	cacheServer *Cache
	once        sync.Once
)

// GetCacheInsOr return cache server instance with given factory.
func GetCacheInsOr(store store.Factory) (*Cache, error) {
	if store != nil {
		once.Do(func() {
			cacheServer = &Cache{store}
		})
	}

	if cacheServer == nil {
		return nil, fmt.Errorf("got nil cache server")
	}

	return cacheServer, nil
}
