// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/stretchr/testify/mock"
)

// MockSrv is a mock implementation of the MockSrv interface.
type MockSrv struct {
	mock.Mock
}

// Create is a mock implementation of the Create method of the MockSrv interface.
func (m *MockSrv) Create(ctx context.Context, item *v1.Item, opts metav1.CreateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

// Update is a mock implementation of the Update method of the MockSrv interface.
func (m *MockSrv) Update(ctx context.Context, item *v1.Item, opts metav1.UpdateOptions) error {
	args := m.Called(ctx, item, opts)
	return args.Error(0)
}

// Delete is a mock implementation of the Delete method of the MockSrv interface.
func (m *MockSrv) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	args := m.Called(ctx, id, opts)
	return args.Error(0)
}

// Get is a mock implementation of the Get method of the MockSrv interface.
func (m *MockSrv) Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.Item, error) {
	args := m.Called(ctx, id, opts)
	return args.Get(0).(*v1.Item), args.Error(1)
}

// List is a mock implementation of the List method of the MockSrv interface.
func (m *MockSrv) List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*v1.ItemList), args.Error(1)
}

// GetAllItems is a mock implementation of the GetAllItems method of the MockSrv interface.
func (m *MockSrv) GetAllItems(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*v1.ItemList), args.Error(1)
}
