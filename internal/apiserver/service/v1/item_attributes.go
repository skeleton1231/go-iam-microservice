package v1

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
)

// ItemAttributeSrv defines functions used to handle item attribute requests.
type ItemAttributeSrv interface {
	Create(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.CreateOptions) error
	Update(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error
	Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.ItemAttributes, error)
	// List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemAttributeList, error)
}

type itemAttributeService struct {
	store store.Factory
}

var _ ItemAttributeSrv = (*itemAttributeService)(nil)

func newItemAttributes(srv *service) *itemAttributeService {
	return &itemAttributeService{store: srv.store}
}

// Implement the methods for ItemAttributeSrv interface here
// Create creates a new item attribute in the storage.
func (i *itemAttributeService) Create(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.CreateOptions) error {
	if err := i.store.ItemAttributes().Create(ctx, itemAttribute, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Update updates an existing item attribute in the storage.
func (i *itemAttributeService) Update(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.UpdateOptions) error {
	if err := i.store.ItemAttributes().Update(ctx, itemAttribute, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Delete deletes an item attribute from the storage.
func (i *itemAttributeService) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	if err := i.store.ItemAttributes().Delete(ctx, id, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Get retrieves an item attribute from the storage.
func (i *itemAttributeService) Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.ItemAttributes, error) {
	itemAttribute, err := i.store.ItemAttributes().Get(ctx, id, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return itemAttribute, nil
}

// List retrieves a list of item attributes from the storage.
// func (i *itemAttributeService) List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemAttributeList, error) {
// 	itemAttributes, err := i.store.ItemAttributes().List(ctx, opts)
// 	if err != nil {
// 		return nil, errors.WithCode(code.ErrDatabase, err.Error())
// 	}

// 	return itemAttributes, nil
// }
