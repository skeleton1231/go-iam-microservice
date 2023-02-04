package authorizer

import (
	"github.com/ory/ladon"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/authzserver/authorization"
)

// PolicyGetter defines function to get policy for a given user.
type PolicyGetter interface {
	GetPolicy(key string) ([]*ladon.DefaultPolicy, error)
}

// Authorization implements authorization.AuthorizationInterface interface.
type Authorization struct {
	getter PolicyGetter
}

// NewAuthorization create a new Authorization instance.
func NewAuthorization(getter PolicyGetter) authorization.AuthorizationInterface {
	return &Authorization{getter}
}
