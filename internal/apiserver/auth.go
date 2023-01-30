// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/spf13/viper"

	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/apiserver/store"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/middleware"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/middleware/auth"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "iam.api.marmotedu.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "iam-apiserver"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"required,username"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {

	return auth.NewBasicStrategy(func(username string, password string) bool {
		// fetch user from database
		user, err := store.Client().Users().Get(context.TODO(), username, metav1.GetOptions{})
		if err != nil {
			return false
		}

		// Compare the login password with the user password.
		if err := user.Compare(password); err != nil {
			return false
		}

		user.LoginedAt = time.Now()
		_ = store.Client().Users().Update(context.TODO(), user, metav1.UpdateOptions{})

		return true
	})
}

func newJWTAuth() middleware.AuthStrategy {
	ginjwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator:    authenticator(),
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refreshResponse(),
		PayloadFunc:     payloadFunc(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return claims[jwt.IdentityKey]
		},
		IdentityKey:  middleware.UsernameKey,
		Authorizator: authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
		// TODO: HTTPStatusMessageFunc:
	})

	return auth.NewJWTStrategy(*ginjwt)
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(string); ok {
			log.L(c).Infof("user `%s` is authenticated.", v)

			return true
		}

		return false
	}
}
