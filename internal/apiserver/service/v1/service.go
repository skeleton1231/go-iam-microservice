// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/service/v1 -destination mock_service.go -package v1 github.com/marmotedu/iam/internal/apiserver/service/v1 Service,UserSrv,SecretSrv,PolicySrv

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
	Items() ItemSrv
	ItemAttributes() ItemAttributesSrv
	ItemImage() ItemImageSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}

func (s *service) Policies() PolicySrv {
	return newPolicies(s)
}

func (s *service) Items() ItemSrv {
	return newItems(s)
}

func (s *service) ItemAttributes() ItemAttributesSrv {
	return newItemAttributes(s)
}

func (s *service) ItemImage() ItemImageSrv {
	return newItemImages(s)
}
