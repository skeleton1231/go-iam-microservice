package item

import (
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/stretchr/testify/mock"
)

type ItemStoreMock struct {
	mock.Mock
}

func (m *ItemStoreMock) Create(item *model.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *ItemStoreMock) Update(item *model.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *ItemStoreMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ItemStoreMock) Get(id int) (*model.Item, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Item), args.Error(1)
}

func (m *ItemStoreMock) List() ([]*model.Item, error) {
	args := m.Called()
	return args.Get(0).([]*model.Item), args.Error(1)
}

)

func TestCreateItemService(t *testing.T) {
	itemStoreMock := new(ItemStoreMock)
	storeFactory := store.NewFactory(itemStoreMock)
	itemService := v1.NewItemService(storeFactory)

	item := &model.Item{
		ASIN: "B123456",
		// Add other fields here
	}

	itemStoreMock.On("Create", item).Return(nil)
	err := itemService.Create(item)
	assert.NoError(t, err)
}

func TestUpdateItemService(t *testing.T) {
	itemStoreMock := new(ItemStoreMock)
	storeFactory := store.NewFactory(itemStoreMock)
	itemService := v1.NewItemService(storeFactory)

	item := &model.Item{
		ASIN: "B123456",
		// Add other fields here
	}

	itemStoreMock.On("Update", item).Return(nil)
	err := itemService.Update(item)
	assert.NoError(t, err)
}

func TestDeleteItemService(t *testing.T) {
	itemStoreMock := new(ItemStoreMock)
	storeFactory := store.NewFactory(itemStoreMock)
	itemService := v1.NewItemService(storeFactory)

	itemID := 1

	itemStoreMock.On("Delete", itemID).Return(nil)
	err := itemService.Delete(itemID)
	assert.NoError(t, err)
}

func TestGetItemService(t *testing.T) {
	itemStoreMock := new(ItemStoreMock)
	storeFactory := store.NewFactory(itemStoreMock)
	itemService := v1.NewItemService(storeFactory)

	itemID := 1
	expectedItem := &model.Item{
		ASIN: "B123456",
		// Add other fields here
	}

	itemStoreMock.On("Get", itemID).Return(expectedItem, nil)
	item, err := itemService.Get(itemID)
	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
}

func TestListItemService(t *testing.T) {
	itemStoreMock := new(ItemStoreMock)
	storeFactory := store.NewFactory(itemStoreMock)
	itemService := v1.NewItemService(storeFactory)

	expectedItems := []*model.Item{
		{
			ASIN: "B123456",
			// Add other fields here
		},
		{
			ASIN: "B789012",
			// Add other fields here
		},
	}

	itemStoreMock.On("List").Return(expectedItems, nil)
	items, err := itemService.List()
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}

