// Copyright 2023 Tal Huang <talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package item

import (
	"context"
	"testing"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestCreateItemService(t *testing.T) {
	mockItemStore := new(MockItemStore)
	storeFactory := store.NewFactory(mockItemStore)
	itemService := v1.NewService(storeFactory)

	item := &model.Item{
		ASIN: "B123456",
		// Add other fields here
	}

	mockItemStore.On("Create", mock.Anything, item, mock.Anything).Return(nil)
	err := itemService.Items().Create(context.Background(), item, metav1.CreateOptions{})
	assert.NoError(t, err)
}
