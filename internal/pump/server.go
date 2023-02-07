// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pump

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pump/pumps"
)

var pmps []pumps.Pump

type pumpServer struct {
	secInterval    int
	omitDetails    bool
	mutex          *redsync.Mutex
	analyticsStore storage.AnalyticsStorage
	pumps          map[string]options.PumpConfig
}

// preparedGenericAPIServer is a private wrapper that enforces a call of PrepareRun() before Run can be invoked.
type preparedPumpServer struct {
	*pumpServer
}
