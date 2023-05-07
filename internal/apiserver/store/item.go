package store

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
)

// ItemStore defines the methods for working with items and their related details in the datastore.
type ItemStore interface {
	Create(ctx context.Context, item *v1.Item, opts metav1.CreateOptions) error
	Update(ctx context.Context, item *v1.Item, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error
	Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.Item, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ItemList, error)
}
