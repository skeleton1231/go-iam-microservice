package v1

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

// ItemSrv defines functions used to handle item requests.
type ItemSrv interface {
	Create(ctx context.Context, item *v1.Item, opts metav1.CreateOptions) error
	Update(ctx context.Context, item *v1.Item, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error
	// DeleteCollection(ctx context.Context, ids []int, opts metav1.DeleteOptions) error
	Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.Item, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error)
}

type itemService struct {
	store store.Factory
}

var _ ItemSrv = (*itemService)(nil)

func newItems(srv *service) *itemService {
	return &itemService{store: srv.store}
}

// Implement the methods for ItemSrv interface here
// Create creates a new item in the storage.
func (i *itemService) Create(ctx context.Context, item *v1.Item, opts metav1.CreateOptions) error {
	if err := i.store.Items().Create(ctx, item, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Update updates an existing item in the storage.
func (i *itemService) Update(ctx context.Context, item *v1.Item, opts metav1.UpdateOptions) error {
	if err := i.store.Items().Update(ctx, item, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Delete deletes an item from the storage.
func (i *itemService) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	if err := i.store.Items().Delete(ctx, id, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Get retrieves an item from the storage.
func (i *itemService) Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.Item, error) {
	item, err := i.store.Items().Get(ctx, id, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return item, nil
}

// List retrieves a list of items from the storage.
func (i *itemService) List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error) {
	items, err := i.store.Items().List(ctx, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return items, nil
}
