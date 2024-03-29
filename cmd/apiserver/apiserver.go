// Copyright 2023 Talhuang <talhuang1231@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// apiserver is the api server for iam-apiserver service.
// it is responsible for serving the platform RESTful resource management.
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver"
)

func main() {

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	// rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}
