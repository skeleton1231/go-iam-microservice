// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/*
mockFactory := new(MockFactory)
mockItemStore := new(MockItemStore)
mockFactory.On("Item").Return(mockItemStore)
mockItemStore.On("Create", ctx, item, opts).Return(nil)

err := mockFactory.Item().Create(ctx, item, opts)
assert.NoError(t, err)
*/
package store

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/stretchr/testify/mock"
)

type MockFactory struct {
	mock.Mock
}

func (m *MockFactory) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFactory) Items() ItemStore {
	args := m.Called()
	return args.Get(0).(ItemStore)
}

func (m *MockFactory) Users() UserStore {
	args := m.Called()
	return args.Get(0).(UserStore)
}

func (m *MockFactory) Secrets() SecretStore {
	args := m.Called()
	return args.Get(0).(SecretStore)
}

func (m *MockFactory) Policies() PolicyStore {
	args := m.Called()
	return args.Get(0).(PolicyStore)
}

func (m *MockFactory) PolicyAudits() PolicyAuditStore {
	args := m.Called()
	return args.Get(0).(PolicyAuditStore)
}

type MockItemStore struct {
	mock.Mock
}

func (m *MockItemStore) Create(ctx context.Context, item *model.Item, opts metav1.CreateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

func (m *MockItemStore) Update(ctx context.Context, item *model.Item, opts metav1.UpdateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

func (m *MockItemStore) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	args := m.Called(ctx, id, opts)
	return args.Error(0)
}

func (m *MockItemStore) Get(ctx context.Context, id int, opts metav1.GetOptions) (*model.Item, error) {
	args := m.Called(ctx, id, opts)
	return args.Get(0).(*model.Item), args.Error(1)
}

func (m *MockItemStore) List(ctx context.Context, opts metav1.ListOptions) (*model.ItemList, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*model.ItemList), args.Error(1)
}
