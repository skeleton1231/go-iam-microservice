// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/*
mockService := new(MockService)
mockItemService := new(MockItemService)
mockService.On("Item").Return(mockItemService)
mockItemService.On("Create", ctx, item, opts).Return(nil)

err := mockService.Item().Create(ctx, item, opts)
assert.NoError(t, err)
*/
package v1

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Item() itemService {
	args := m.Called()
	return args.Get(0).(itemService)
}

type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) Create(ctx context.Context, item *model.Item, opts metav1.CreateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

func (m *MockItemService) Update(ctx context.Context, item *model.Item, opts metav1.UpdateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

func (m *MockItemService) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	args := m.Called(ctx, id, opts)
	return args.Error(0)
}

func (m *MockItemService) Get(ctx context.Context, id int, opts metav1.GetOptions) (*model.Item, error) {
	args := m.Called(ctx, id, opts)
	return args.Get(0).(*model.Item), args.Error(1)
}

func (m *MockItemService) List(ctx context.Context, opts metav1.ListOptions) (*model.ItemList, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*model.ItemList), args.Error(1)
}
